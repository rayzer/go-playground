package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type valueElement struct {
	key   int
	value int
}

type LRUCache struct {
	dict   map[int]*list.Element
	list   *list.List
	remain int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dict:   make(map[int]*list.Element, capacity),
		list:   list.New(),
		remain: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.dict[key]; ok {
		this.list.MoveToFront(element)
		return element.Value.(valueElement).value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if element, ok := this.dict[key]; ok {
		element.Value = valueElement{key, value}
		this.list.MoveToFront(element)
	} else {
		if this.remain > 0 {
			this.remain--
		} else {
			lastElement := this.list.Back()
			delete(this.dict, lastElement.Value.(valueElement).key)
			this.list.Remove(lastElement)

		}
		this.dict[key] = this.list.PushFront(valueElement{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {
	//lRUCache := Constructor(2)
	//lRUCache.Put(1, 1)           // cache is {1=1}
	//lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	//fmt.Println(lRUCache.Get(1)) // return 1
	//lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	//fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	//lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	//fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	//fmt.Println(lRUCache.Get(3)) // return 3
	//fmt.Println(lRUCache.Get(4)) // return 4
	lRUCache := Constructor(1)
	lRUCache.Put(2, 1)
	fmt.Println(lRUCache.Get(2))
}
