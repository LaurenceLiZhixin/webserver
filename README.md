## 开发 web 服务程序

### 1、概述

开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

**任务目标**

1. 熟悉 go 服务器工作原理
2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
3. 使用 curl 工具访问 web 程序
4. 对 web 执行压力测试

### 2、任务要求

1. 编程 web 服务程序 类似 cloudgo 应用。
   - 要求有详细的注释
   - 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

### 3、 任务实现

1. 开启服务

   ```
   server := service.NewServer()
   server.Run(":" + port)
   ```

2. 构建使用json格式的Server，并使用路由器mx

   ```
   func NewServer() *negroni.Negroni {
   	formatter := render.New(render.Options{
   		IndentJSON: true,
   	})
   	n := negroni.Classic()
   	mx := mux.NewRouter()
   	initRoutes(mx, formatter)
   	n.UseHandler(mx)
   	return n
   }
   ```

3. 为当前路由器增加路由

   ```
   mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
   mx.HandleFunc("/add/{id1}/{id2}", addHandler).Methods("GET")
   ```

   可看到增加了两条GET方式的路由，其中add路径为自定义的web服务，可以计算两个参数的数值之和，并返还给客户端。

4. 具体addHandler的实现：

   ```
   func addHandler(w http.ResponseWriter, req *http.Request) {
   	formatter := render.New(render.Options{
   		IndentJSON: true,
   	})
   	vars := mux.Vars(req)
   	id1 := vars["id1"]
   	id2 := vars["id2"]
   	id_1, _ := strconv.Atoi(id1)
   	id_2, _ := strconv.Atoi(id2)
   	id := id_1 + id_2
   	id_str := strconv.Itoa(id)
   	formatter.JSON(w, http.StatusOK, struct{ Test string }{id1 + " + " + id2 + " = " + id_str})
   }
   ```

   这个函数从url参数中获取到两个数字，并进行计算，最终按照json格式返还给客户端。

### 4、 测试

#### 1. 使用curl进行测试

![微信截图_20191106203150](img\微信截图_20191106203150.png)

可看到收到回复结果 “123 + 456 = 579”

#### 2. 使用ab进行压力测试

到官网下载ApacheHaus

https://www.apachehaus.com/cgi-bin/download.plx?dli=lVUMpFWVBVTT6N2aiVEezokVOpkVFVVcT1GZTJVQ

下载后在bin目录下进行测试。

![1573044490567](img\1573044490567.png)

逐字段分析

Benchmarking localhost (be patient)
Completed 100 requests //进行并发压力测试 参数中要求进行1000次测试
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software://服务软件
Server Hostname:        localhost //主机名
Server Port:            8080 //端口

Document Path:          /add/123/456
Document Length:        32 bytes

Concurrency Level:      100 // 并发级别
Time taken for tests:   1.329 seconds //总共测试多长时间
Complete requests:      1000 //并发数
Failed requests:        0 //失效请求
Total transferred:      155000 bytes //总共传输字节数
HTML transferred:       32000 bytes // HTML字节数，实际的页面传递的字节数
Requests per second:    752.19 [#/sec] (mean) // 每秒多少个请求，代表服务器的吞吐量
Time per request:       132.945 [ms] (mean) // 用户平均的等待时间

// 服务器的平均处理时间

Time per request:       1.329 [ms] (mean, across all concurrent requests) 
Transfer rate:          113.86 [Kbytes/sec] received //每秒获取的数据长度

Connection Times (ms) //链接的时间
               min  mean[+/-sd] median   max // 最短时间 平均时间 中值 最长时间
Connect:        0    0   0.3      0       1  //链接时间
Processing:     2  125  24.7    131     250//处理时间
Waiting:        2  125  24.8    131     250 //等待时间
Total:          2  125  24.7    131     250 //合计总时间

Percentage of the requests served within a certain time (ms)
  50%    131 //一半的请求在131ms内返回
  66%    134 //66%的请求在134ms内返回
  75%    136
  80%    137
  90%    138
  95%    142
  98%    145
  99%    148 
 100%    250 (longest request)