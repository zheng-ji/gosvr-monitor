### gosvr-monitor

#### 简介

检测程序运行的各项数据，Goroutine 堆栈等

### 编译和运行

```
cd example
go build test_monit.go
./test_monit
```

### 访问

```
http://127.0.0.1:6062
```

### 实现介绍

```
curl "http://127.0.0.1:6062/info"

Cache HitRate: 0.742956
QPS: 47
Goroutine Num: 28
PersisWrite(ms): 0.000000
PersisRead(ms): 3.410000
InternalCacheRead (ms): 0.000000
clustercacheRead (ms): 1.700000
Memstats Alloc: 11512568
MemStats TotalAlloc: 43079377144
MemStats Sys: 38168824
MemStats Lookups: 1629304
MemStats Mallocs: 544423474
MemStats Frees: 544365072
MemStats HeepAlloc: 11512568
MemStats HeepSys: 28311552
MemStats HeepIdle: 13746176
MemStats HeepInuse: 14565376
MemStats HeepReleased: 0
MemStats HeepObjects: 58402
MemStats NextGC: 16033600
MemStats LastGC: 1439188926763571035
MemStats PauseTotalNs: 3855139783
MemStats NumGC: 4938
```

