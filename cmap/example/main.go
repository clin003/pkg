package main

import (
	"fmt"

	"gitee.com/lyhuilin/pkg/cmap"
)

var KVcmap cmap.ConcurrentMap

//小写init可以自动执行
func init() {
	var concurrency int
	concurrency = 10
	// var pairRedistributor cmap.PairRedistributor
	KVcmap, _ = cmap.NewConcurrentMap(concurrency, nil)
}

func main() {
	if ok, err := KVcmap.Put("Hi", "World"); !ok || err != nil {
		return
	}
	v := KVcmap.Get("Hi")
	fmt.Println(v)
}
