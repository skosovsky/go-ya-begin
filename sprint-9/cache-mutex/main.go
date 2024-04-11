package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[int]int
}

func (c *Cache) Get(id int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.data[id]
	if ok {
		return v
	}
	// получаем значение для указанного ключа
	v = 2 * id
	c.data[id] = v
	return v
}

func main() {
	cache := Cache{
		data: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(cache.data))
}
