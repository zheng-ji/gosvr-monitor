/*
Copyright (c) 2015, zheng-ji.info
*/

package monitor

import (
	"fmt"
	"github.com/zheng-ji/gosvr-monitor"
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	fmt.Println("start")
	monitor.InitMonitor([]string{"WRITE", "READ"}, 1)
	monitor.StartMonitorServer("0.0.0.0:7070")
	timeStart := time.Now()
	time.Sleep(300 * time.Millisecond)
	monitor.StatByAction("READ", timeStart)
	fmt.Println("end")
}
