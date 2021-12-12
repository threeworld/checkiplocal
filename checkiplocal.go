package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gosuri/uitable"
	flags "github.com/jessevdk/go-flags"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var opts struct {
	Threads int  `short:"t" long:"threads" default:"8" description:"How many threads should be used"`
	City    bool `short:"c" long:"city" description:"Print the city where the IP is located"`
	Region  bool `short:"r" long:"region" description:"Print the area where the IP is located"`
	Addr    bool `short:"a" long:"addr" description:"Print the exact address of the IP by default"`
}

var sm sync.Map

func main() {

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil || !(opts.Addr || opts.City || opts.Region) {
		os.Exit(1)
	}
	work := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			work <- s.Text()
		}
		if err := s.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		close(work)
	}()
	wg := &sync.WaitGroup{}
	for i := 0; i < opts.Threads; i++ {
		wg.Add(1)
		go doWork(work, wg)
	}
	wg.Wait()

	table := uitable.New()
	table.MaxColWidth = 80
	table.AddRow("IP", "ADDR", "CITY", "REGION")
	sm.Range(func(key, value interface{}) bool {
		ipInfo := value.(map[string]string)
		table.AddRow(ipInfo["ip"], ipInfo["addr"], ipInfo["city"], ipInfo["region"])
		return true
	})
	fmt.Println(table)
}

func doWork(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range work {
		ip := net.ParseIP(line)
		if ip == nil {
			addResults(line, "")
		} else {
			ret, _ := getIPLocal(ip.String())
			addResults(line, ret)
		}
	}
}

func addResults(ip, result string) {
	if result == "" {
		sm.Store(ip, map[string]string{"ip": ip})
	} else {
		ipInfo := make(map[string]string)
		err := json.Unmarshal([]byte(result), &ipInfo)
		// 转换失败
		if err != nil {
			sm.Store(ip, map[string]string{"ip": ip})
		} else {

			if !opts.Addr {
				ipInfo["addr"] = ""
			}
			if !opts.City {
				ipInfo["city"] = ""
			}
			if !opts.Region {
				ipInfo["region"] = ""
			}
			ipInfo["addr"] = strings.TrimSpace(ipInfo["addr"])
			sm.Store(ip, ipInfo)
		}
	}
}

func getIPLocal(ip string) (string, error) {
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip="
	req, err := http.NewRequest("GET", url+ip, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//关闭resp,避免内存泄露
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	//转换编码 gbk转utf-8
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	ret, _ := ioutil.ReadAll(reader)
	return string(ret), nil
}
