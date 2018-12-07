package main

import "sync"

// 类型以字段形式直接嵌入
type myLocker struct {
	sync.Mutex
}
type myLocker2 sync.Locker

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}
