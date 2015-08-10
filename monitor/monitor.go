// Author: zheng-ji.info

package monitor

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type StatUnit struct {
	sumMs        int64
	operateTime  int
	operateAvgMs float64
	operateLock  *sync.Mutex
}

type monitor struct {
	//各类统计信息
	statMaps map[int]*StatUnit
}

var defaultMonitor *monitor

func init() {
	defaultMonitor = &monitor{
		statMaps: make(map[int]*StatUnit),
	}

	for i := 0; i < STAT_UNINT_NUM; i++ {
		defaultMonitor.statMaps[i] = &StatUnit{
			sumMs:        0,
			operateTime:  0,
			operateAvgMs: 0,
			operateLock:  new(sync.Mutex),
		}
	}
}

// 各种类型时长统计
func StatByAction(statType int, t time.Time) {

	defaultMonitor.statMaps[statType].operateLock.Lock()

	defaultMonitor.statMaps[statType].sumMs += int64(Since(t))
	defaultMonitor.statMaps[statType].operateTime++

	if defaultMonitor.statMaps[statType].operateTime >= NUM_PER_STAT {
		defaultMonitor.statMaps[statType].operateAvgMs = float64(defaultMonitor.statMaps[statType].sumMs) / float64(defaultMonitor.statMaps[statType].operateTime)
		defaultMonitor.statMaps[statType].operateTime = 0
		defaultMonitor.statMaps[statType].sumMs = 0
	}

	defaultMonitor.statMaps[statType].operateLock.Unlock()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	memstats := new(runtime.MemStats)
	runtime.ReadMemStats(memstats)

	fmt.Fprintf(w,
		"Goroutine Num: %d\n"+
			"PersisWrite(ms): %f\n"+
			"PersisRead(ms): %f\n"+
			"InternalCacheRead (ms): %f\n"+
			"clustercacheRead (ms): %f\n"+
			"Memstats Alloc: %d\n"+
			"MemStats TotalAlloc: %d\n"+
			"MemStats Sys: %d\n"+
			"MemStats Lookups: %d\n"+
			"MemStats Mallocs: %d\n"+
			"MemStats Frees: %d\n"+
			"MemStats HeepAlloc: %d\n"+
			"MemStats HeepSys: %d\n"+
			"MemStats HeepIdle: %d\n"+
			"MemStats HeepInuse: %d\n"+
			"MemStats HeepReleased: %d\n"+
			"MemStats HeepObjects: %d\n"+
			"MemStats NextGC: %d\n"+
			"MemStats LastGC: %d\n"+
			"MemStats PauseTotalNs: %d\n"+
			"MemStats NumGC: %d\n",
		runtime.NumGoroutine(),
		defaultMonitor.statMaps[STAT_PERSIS_WRITE].operateAvgMs,
		defaultMonitor.statMaps[STAT_PERSIS_READ].operateAvgMs,
		defaultMonitor.statMaps[STAT_INNERCACHE_READ].operateAvgMs,
		defaultMonitor.statMaps[STAT_CACHECLUSTER_READ].operateAvgMs,
		memstats.Alloc,
		memstats.TotalAlloc,
		memstats.Sys,
		memstats.Lookups,
		memstats.Mallocs,
		memstats.Frees,
		memstats.HeapAlloc,
		memstats.HeapSys,
		memstats.HeapIdle,
		memstats.HeapInuse,
		memstats.HeapReleased,
		memstats.HeapObjects,
		memstats.NextGC,
		memstats.LastGC,
		memstats.PauseTotalNs,
		memstats.NumGC)
}

// 启动http服务监听
func StartMonitorServer(addr string) {
	fmt.Println("start")

	go func() {
		http.HandleFunc("/info", infoHandler)
		http.ListenAndServe(addr, nil)
	}()
}
