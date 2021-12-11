# checkiplocal
go实现的快速批量查询IP归属地小工具

## 简单用法

单个IP
```
echo 223.104.246.67 | ./checkiplocal -a -r -c
```

多个IP
```
cat ips.txt |./checkiplocal -a -r -c
```

## 命令行选项

```
Usage:
  checkiplocal [OPTIONS]

Application Options:
  -t, --threads= How many threads should be used (default: 8)
  -c, --city     Print the city where the IP is located
  -r, --region   Print the area where the IP is located
  -a, --addr     Print the exact address of the IP by default

Help Options:
  -h, --help     Show this help message
```
  
 
