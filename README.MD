### Router Manage DB lite —— 轻量级的路由器管理数据库

#### 安装
##### 编译安装
依赖：
golang 1.8+
gcc

```
# git clone https://github.com/shanghai-edu/rmdb-lite.git
cd rmdb-lite
go get ./...
./control build
```
##### 二进制安装

#### 运行
```
# ./control -h
./control build|pack|start|stop|restart|status|tail
```
#### 接口

##### 单个 ip 查询
role = 1
```
# curl -H "X-API-KEY: 3a1179104ae0188ed4a3c100b38343ff" http://127.0.0.1:18080/api/v1/router?ip=192.25.209.100   
                                                                                                        
{"code":0,"message":"ok","version":"0.1","data":{"ip":"192.25.209.100","node":"西宁","node_detail":""}}
```

role = 2
```
# curl -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router?ip=192.25.209.100

{"code":0,"message":"ok","version":"0.1","data":{"ip":"192.25.209.100","node":"西宁","node_detail":"北京街I座"}}
```
##### 多个 ip 查询
role = 1
```
# curl -X POST -d "{\"ips\":[\"192.25.209.100\",\"192.52.207.167\",\"192.76.5.75\"]}" -H "Content-type: application/json" -H "X-API-KEY: 3a1179104ae0188ed4a3c100b38343ff" http://127.0.0.1:18080/api/v1/router/multi    

{"code":0,"message":"ok","version":"0.1","data":{"routers":[{"ip":"192.25.209.100","node":"西宁","node_detail":""},{"ip":"192.52.207.167","node":"天津","node_detail":""},{"ip":"192.76.5.75","node":"昆明","node_detail":""}],"failed_list":[]}}
```
role = 2
```
# curl -X POST -d "{\"ips\":[\"192.25.209.100\",\"192.52.207.167\",\"192.76.5.75\"]}" -H "Content-type: application/json" -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router/multi

{"code":0,"message":"ok","version":"0.1","data":{"routers":[{"ip":"192.25.209.100","node":"西宁","node_detail":"北京街I座"},{"ip":"192.52.207.167","node":"天津","node_detail":"李路n座"},{"ip":"192.76.5.75","node":"昆明","node_detail":"李路d座"}],"failed_list":[]}}
```
##### 读取所有节点
需要 role = 3 权限
```
# curl -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router/all

…………
```
##### 删除节点
需要 role =3 权限
```
# curl -X DELETE -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router?ip=192.25.209.100                                

{"code":0,"message":"ok","version":"0.1","data":null}
```
##### 创建节点
需要 role =3 权限
```
# curl -X POST -d "{\"ip\":\"192.25.209.100\",\"node\": \"西宁\",\"node_detail\":\"北京街I座\"}" -H "Content-type: application/json" -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router  
 
{"code":0,"message":"ok","version":"0.1","data":null}  
```
##### 更新节点
需要 role = 3 权限
```
# curl -X PUT -d "{\"ip\":\"192.25.209.100\",\"node\": \"西宁\",\"node_detail\":\"北京街I座\"}" -H "Content-type: application/json" -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" http://127.0.0.1:18080/api/v1/router   

{"code":0,"message":"ok","version":"0.1","data":null}
```
#### 性能
##### 查询性能
```
# ./wrk -t12 -c400 -d30s -H "X-API-KEY: 3a1179104ae0188ed4a3c100b38343ff" http://127.0.0.1:18080/api/v1/router?ip=192.25.209.100 
Running 30s test @ http://127.0.0.1:18080/api/v1/router?ip=192.25.209.100
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    46.58ms   53.37ms 387.52ms   81.42%
    Req/Sec     1.14k   245.79     2.18k    73.15%
  408803 requests in 30.05s, 208.58MB read
Requests/sec:  13604.58
Transfer/sec:      6.94MB
```
##### 写入性能
```
# ./wrk -t12 -c400 -d30s -H "X-API-KEY: 771261d709af69da179a330ed8f16fa1" -s put.lua http://127.0.0.1:18080/api/v1/router   
Running 30s test @ http://127.0.0.1:18080/api/v1/router
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   510.12ms  592.93ms   2.00s    80.21%
    Req/Sec    16.38     12.78   121.00     74.50%
  4661 requests in 30.03s, 2.18MB read
  Socket errors: connect 0, read 0, write 0, timeout 2160
  Non-2xx or 3xx responses: 908
Requests/sec:    155.19
Transfer/sec:     74.26KB
```