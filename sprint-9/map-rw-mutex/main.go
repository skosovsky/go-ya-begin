package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	mu   sync.RWMutex
	data map[string]string
}

func (m *Map) Get(key string) string {
	// RLock() даёт возможность нескольким горутинам читать данные
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data[key]
}

func (m *Map) Set(key string, v string) {
	// Lock() блокирует все операции
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = v
}

func main() {
	cacheData := Map{
		mu:   sync.RWMutex{},
		data: make(map[string]string),
	}
	for i := 0; i < 10; i++ {
		go func() {
			for {
				cacheData.Set("a", ".")
				time.Sleep(50 * time.Millisecond) //nolint:mnd // it's learning code
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			count := 0
			for {
				v := cacheData.Get("a")
				count++
				if count%10 == 0 {
					fmt.Print(v) //nolint:forbidigo // it's learning code
				}
				time.Sleep(20 * time.Millisecond) //nolint:mnd // it's learning code
			}
		}()
	}
	time.Sleep(1 * time.Second)
}
