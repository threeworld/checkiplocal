# checkiplocal
go实现的快速批量查询IP归属地小工具

## 简单用法
将 IP 地址列表通过管道传输到工具中:
单个IP
```
echo 223.104.246.67 | ./checkiplocal -a -r -c
```

多个IP
```
cat ips.txt |./checkiplocal -a -r -c
```

输出
```
IP              ADDR                                    CITY    REGION
1.13.189.46     吉林省长春市 方正宽带                   长春市        
1.117.204.157   北京市 中电华通                         北京市        
1.117.199.237   北京市 中电华通                         北京市        
1.116.179.58    北京市 中电华通                         北京市        
1.116.41.177    北京市 中电华通                         北京市        
1.116.140.20    北京市 中电华通                         北京市        
1.12.221.250    北京市 北京北大方正宽带网络科技有限公司 北京市        
1.12.236.147    北京市 北京北大方正宽带网络科技有限公司 北京市        
1.117.229.146   北京市 中电华通                         北京市        
1.117.176.186   北京市 中电华通                         北京市        
1.117.189.110   北京市 中电华通                         北京市        
1.117.81.82     北京市 中电华通                         北京市        
1.117.248.245   北京市 中电华通                         北京市        
1.117.204.147   北京市 中电华通                         北京市  
....
```
## 命令行选项

```
Usage:
  checkiplocal [OPTIONS]

Application Options:
  -t, --threads  How many threads should be used (default: 8)
  -c, --city     Print the city where the IP is located
  -r, --region   Print the area where the IP is located
  -a, --addr     Print the exact address of the IP by default

Help Options:
  -h, --help     Show this help message
```
  
 
