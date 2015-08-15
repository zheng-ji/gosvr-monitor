// Author: zheng-ji.info

package main

import (
	"fmt"
	"github.com/zheng-ji/gosvr-monitor"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func monitor_write() {
	timeStart := time.Now()
	time.Sleep(500 * time.Millisecond)
	defer func() {
		//defer 的时候统计监控, 用goroutine 使得不影响性能
		go monitor.StatByAction("WRITE", timeStart)
	}()
}

func monitor_read() {
	timeStart := time.Now()
	time.Sleep(300 * time.Millisecond)
	defer func() {
		//defer 的时候统计监控, 用goroutine 使得不影响性能
		go monitor.StatByAction("READ", timeStart)
	}()
}

func main() {

	//初始化monitor, 自定义监控的命令,如READ,WRITE, 并启动监控服务
	monitor.InitMonitor([]string{"WRITE", "READ"}, 1)
	monitor.StartMonitorServer("0.0.0.0:7070")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGINT)

	monitor_read()
	monitor_write()

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
