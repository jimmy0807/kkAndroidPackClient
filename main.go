package main

import (
	"fmt"
	"kkAndroidPackClient/config"
	httpServer "kkAndroidPackClient/http"
	"kkAndroidPackClient/packManager"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	packManager.Instance()
	var stopLock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		stopLock.Lock()
		stop = true
		stopLock.Unlock()
		fmt.Println("程序要退出了")
		log.Println("Cleaning before stop...")
		stopChan <- struct{}{}
		os.Exit(0)
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	handler := httpServer.HandleHTTPServer("/")
	url := config.HttpHost + ":" + config.HttpPort
	httpServer.Start(url, handler)
}
