package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{
		store: sync.Map{},
	}
}

func (c *Cache) Set(key string, value int, d time.Duration) {
	// добавить значение в мапу
	c.store.Store(key, value)

	time.AfterFunc(d, func() {
		c.store.Delete(key)
	})
}

func (c *Cache) Get(key string) (any, bool) {
	// получить и возвратить значение мапы
	return c.store.Load(key)
}

func main() {
	// функция main() не требует изменений

	cache := NewCache()

	cache.Set("1", 567, 100*time.Millisecond) //nolint:mnd // it's learning code
	cache.Set("2", 22, 200*time.Millisecond)  //nolint:mnd // it's learning code
	cache.Set("3", 9, 300*time.Millisecond)   //nolint:mnd // it's learning code

	// определяем локальную функцию
	printData := func() {
		for _, key := range []string{"1", "2", "3"} {
			v, ok := cache.Get(key)
			if ok {
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}

	printData()
	time.Sleep(150 * time.Millisecond)
	printData()
	time.Sleep(100 * time.Millisecond)
	printData()
}
