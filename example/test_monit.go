// Author: zheng-ji.info

package main

import (
	"fmt"
	"gosvr-monitor/monitor"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func test_func1() {
	timeStart := time.Now()
	defer func() {
		go monitor.StatByAction(monitor.STAT_PERSIS_WRITE, timeStart)
	}()
}

func main() {
	// 启动监控服务
	monitor.StartMonitorServer("0.0.0.0:6062")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGINT)

	test_func1()

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
