package main

import (
	"log"
	"time"
)

type taskExecutor struct {
	ticker *time.Ticker
	stop   chan bool
}

const DefaultSyncInterval = 2

var fsTaskExecutor *taskExecutor

func runFSTaskInTicker() {
	fsTaskExecutor = new(taskExecutor)
	fsTaskExecutor.stop = make(chan bool)
	fsTaskExecutor.ticker = time.NewTicker(time.Second * time.Duration(DefaultSyncInterval))
	for {
		select {
		case <-fsTaskExecutor.stop:
			log.Println("executor stop requested")
			return
		case <-fsTaskExecutor.ticker.C:
			fsTaskExecutor.ticker.Stop()
			log.Println("time ticking")
			fsTaskExecutor.ticker = time.NewTicker(time.Second * time.Duration(DefaultSyncInterval))
		}
	}
}

func stopScheduledFSTask() {
	if fsTaskExecutor.stop != nil {
		close(fsTaskExecutor.stop)
	}
	fsTaskExecutor.stop = nil
	log.Println("stopped")
}

func main() {
	go runFSTaskInTicker()
	time.Sleep(time.Second * 20)
	stopScheduledFSTask()
}
