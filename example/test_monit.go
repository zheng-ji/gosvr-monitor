// Copyright 2015
// Author: zheng-ji.info

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/zheng-ji/gosvr-monitor/monitor"
	"time"
)

func test_func1() {
	timeStart := time.Now()
	def func() {
		go monitor.StatByAction(monitor)
	}
}

func main() {
	// 启动监控服务
	monitor.StartMonitorServer("0.0.0.0:6062")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGINT)


L:
	for {
		select {
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGUSR1:
				fmt.Println("Reopen log file")
				// TODO reopen file
			case syscall.SIGTERM, syscall.SIGINT:
				fmt.Println("Catch SIGTERM singal, exit.")
				break L
			}
		}
	}
}
