### gosvr-monitor

#### 简介

检测程序运行的各项数据，Goroutine 堆栈等

### 编译和运行

```
go get github.com/zheng-ji/govr-monitor
```

### 使用范例

```go
import (
    "github.com/zheng-ji/gosvr-monitor"
)
func fun_test() {
	timeStart := time.Now()
	defer func() {
		//defer 的时候统计监控, 用goroutine 使得不影响性能
		go monitor.StatByAction("WRITE", timeStart)
	}()
    ....
}

func main() {
	//初始化monitor, 自定义监控的命令,如READ,WRITE, 以及每次统计的法值
    // 启动监控服务
	monitor.InitMonitor([]string{"WRITE", "READ"}, 1)
	monitor.StartMonitorServer("0.0.0.0:7070")
    fun_test
    ...
}
```

### 访问

```
http://127.0.0.1:7070
```

### 输出

```
curl "http://127.0.0.1:7070/info"
WRITE (ms):500.000000
READ (ms):300.000000

```

