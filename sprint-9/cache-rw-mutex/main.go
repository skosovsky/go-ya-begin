package main

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	m  map[int]int
}

func (cache *Cache) Get(i int) int {
	cache.mu.RLock()
	v, ok := cache.m[i]
	cache.mu.RUnlock()
	if ok {
		return v
	}

	// получаем значение для указанного ключа
	v = 2 * i

	cache.mu.Lock()
	cache.m[i] = v
	cache.mu.Unlock()

	return v
}

func main() {
	cache := Cache{
		m: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}

	time.Sleep(1 * time.Second)
	log.Println(len(cache.m))
}
