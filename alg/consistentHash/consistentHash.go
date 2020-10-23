package main

import (
	"errors"
	//"hash/crc32"
	//"github.com/spaolacci/murmur3"
	"github.com/dgryski/go-farm"
	"sort"
	"sync"
)

type uints []uint32

func (x uints) Len() int           { return len(x) }
func (x uints) Less(i, j int) bool { return x[i] < x[j] }
func (x uints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type server struct {
	address string
}

type servers []server

type ConsistentHash struct {
	nodes        servers
	circle       map[uint32]server //the ring
	servers      map[server]bool
	multiplex    int   //how may vnode for one real server
	count        int64 //current server number
	sortedHashes uints
	sync.RWMutex
}

func New(m int) *ConsistentHash {
	return &ConsistentHash{
		circle:    make(map[uint32]server),
		servers:   make(map[server]bool),
		multiplex: m,
	}
}

func (c *ConsistentHash) AddServer(s server) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.multiplex; i++ {
		c.circle[c.hash(s.address+string(i))] = s
	}
	c.servers[s] = true
	c.updateSortedHashes()
	c.count++
}

func (c *ConsistentHash) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	if cap(c.sortedHashes)/(c.multiplex*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
}

func (c *ConsistentHash) hash(s string) uint32 {
	//if len(s) < 64 {
	//	var scratch [64]byte
	//	copy(scratch[:], s)
	//	return crc32.ChecksumIEEE(scratch[:len(s)])
	//}
	//return crc32.ChecksumIEEE([]byte(s))
	//return murmur3.Sum32([]byte(s))
	return farm.Hash32([]byte((s)))
}

func (c *ConsistentHash) search(key uint32) (idx int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	idx = sort.Search(len(c.sortedHashes), f)
	if idx >= len(c.sortedHashes) {
		idx = 0
	}
	return
}

func (c *ConsistentHash) FindTargetServer(element string) (server, error) {
	c.Lock()
	defer c.Unlock()

	if len(c.circle) == 0 {
		return server{}, errors.New("empty circle")
	}

	hash := c.hash(element)
	serverIdx := c.search(hash)
	return c.circle[c.sortedHashes[serverIdx]], nil
}
