package main

import (
	"fmt"
	"github.com/gonum/stat"
	"sort"
)

func main() {
	data := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta"}
	result := make(map[int]float64)
	for i := 5; i <= 500; i += 5 {
		c := New(i) //vnode for each server
		c.AddServer(server{"cache0.test.com"})
		c.AddServer(server{"cache1.test.com"})
		c.AddServer(server{"cache2.test.com"})
		fmt.Println(len(c.circle))
		hitCount := make(map[string]float64)
		for j := 0; j < 10000; j++ {
			for _, datum := range data {
				s, err := c.FindTargetServer(datum + string(j))
				if err == nil {
					hitCount[s.address]++
				}
			}
		}
		var tmp []float64
		for _, v := range hitCount {
			tmp = append(tmp, v)
		}
		result[i] = stat.StdDev(tmp, nil)
		fmt.Println("vnode number", i, hitCount)
	}

	var keys []int
	for k := range result {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, ",", result[key])
	}
}
