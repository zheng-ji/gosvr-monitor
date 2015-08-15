// Author: zheng-ji.info

package monitor

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

//统计单元
type StatUnit struct {
	sumMs        int64
	operateTime  int
	operateAvgMs float64
	operateLock  *sync.Mutex
}

//统计机
type monitor struct {
	statMaps map[string]*StatUnit
	numLimit int
}

var defaultMonitor *monitor

func InitMonitor(stat_items []string, num_of_limit int) {
	defaultMonitor = &monitor{
		statMaps: make(map[string]*StatUnit),
		numLimit: num_of_limit,
	}

	fmt.Println(len(stat_items))
	for i := 0; i < len(stat_items); i++ {
		defaultMonitor.statMaps[stat_items[i]] = &StatUnit{
			sumMs:        0,
			operateTime:  0,
			operateAvgMs: 0,
			operateLock:  new(sync.Mutex),
		}
	}
}

// 各种类型时长统计,加锁统计
func StatByAction(statType string, t time.Time) {

	defaultMonitor.statMaps[statType].operateLock.Lock()

	defaultMonitor.statMaps[statType].sumMs += int64(Since(t))
	defaultMonitor.statMaps[statType].operateTime++

	if defaultMonitor.statMaps[statType].operateTime >= defaultMonitor.numLimit {
		defaultMonitor.statMaps[statType].operateAvgMs = float64(defaultMonitor.statMaps[statType].sumMs) / float64(defaultMonitor.statMaps[statType].operateTime)
		defaultMonitor.statMaps[statType].operateTime = 0
		defaultMonitor.statMaps[statType].sumMs = 0
	}

	defaultMonitor.statMaps[statType].operateLock.Unlock()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	respstr := ""
	for k, v := range defaultMonitor.statMaps {
		respstr += fmt.Sprintf("%s (ms):%f\n", k, v.operateAvgMs)
	}
	fmt.Fprintf(w, respstr)
}

// 启动http服务监听
func StartMonitorServer(addr string) {
	fmt.Println("start")

	go func() {
		http.HandleFunc("/info", infoHandler)
		http.ListenAndServe(addr, nil)
	}()
}
