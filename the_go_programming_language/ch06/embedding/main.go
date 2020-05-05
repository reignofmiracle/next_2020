package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

var test = struct {
	a int
	b int
}{
	a: 10,
	b: 13,
}

func main() {
	cache.mapping["test1"] = "test1"
	cache.mapping["test2"] = "test2"

	fmt.Println(lookup("test1"))
}

func lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
