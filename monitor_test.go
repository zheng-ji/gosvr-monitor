/*
Copyright (c) 2015, zheng-ji.info
*/

package monitor

import (
	"fmt"
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	InitMonitor([]string{"WRITE", "READ"}, 1)
	StartMonitorServer("0.0.0.0:7070")
	timeStart := time.Now()
	time.Sleep(300 * time.Millisecond)
	StatByAction("READ", timeStart)
}
