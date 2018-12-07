package main

import (
	"log"
	"sync"
	"time"
)

type taskExecutor struct {
	ticker *time.Ticker
	stop   chan bool
}

const DefaultSyncInterval = 500

var fsTaskExecutor *taskExecutor

func runFSTaskInTicker(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	fsTaskExecutor = new(taskExecutor)
	fsTaskExecutor.stop = make(chan bool)
	fsTaskExecutor.ticker = time.NewTicker(time.Millisecond * time.Duration(DefaultSyncInterval))
	defer close(fsTaskExecutor.stop)
	defer log.Println("Task Ticker stopped")

	for {
		select {
		case <-fsTaskExecutor.stop:
			log.Println("get stop chan")
			return
		case <-fsTaskExecutor.ticker.C:
			fsTaskExecutor.ticker.Stop()
			log.Println("time ticking")
			fsTaskExecutor.ticker = time.NewTicker(time.Millisecond * time.Duration(DefaultSyncInterval))
		}
	}

}

func stopScheduledFSTask(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	log.Println("call executor.ticker.stop")
	if fsTaskExecutor.stop != nil {
		log.Println("close executor.stop")
		fsTaskExecutor.stop <- true
	}
	log.Println("stopped")
}

func main() {
	var wg sync.WaitGroup
	go runFSTaskInTicker(&wg)
	time.Sleep(time.Second * 3)
	stopScheduledFSTask(&wg)
	wg.Wait()
}
