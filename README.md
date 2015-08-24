### gosvr-monitor

Customize Your Monitor Items, Inspect the time each of them spend

自定义监控条目，检测程序运行时的各个操作耗费时长

### Complie

```
go get github.com/zheng-ji/gosvr-monitor
```

### Example

```go
import (
    "github.com/zheng-ji/gosvr-monitor"
)
func func_test() {
	timeStart := time.Now()
	defer func() {
		//defer 的时候统计监控, 用goroutine 使得不影响性能
		go monitor.StatByAction("WRITE", timeStart)
	}()
    ....
}

func main() {
	// 初始化monitor, 自定义监控的命令,如READ,WRITE等自定义名称, 以及每次统计的阀值
    // 启动监控服务
	monitor.InitMonitor([]string{"WRITE", "READ"}, 1)
	monitor.StartMonitorServer("0.0.0.0:7070")
    func_test()
    ...
}
```

### How To Use

```
curl "http://127.0.0.1:7070/info"
```

### OutPut

```
curl "http://127.0.0.1:7070/info"
WRITE (ms):15.000000
READ (ms):3.000000
```
------
MIT License
