package test_mem

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

// Queue описывает очередь.
type Queue struct {
	first *QueueItem // указатель на первый элемент очереди
	last  *QueueItem // указатель на последний элемент очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	value string     // данные
	next  *QueueItem // указатель на следующий элемент
}

// PopWithNil удаляет первый элемент из очереди и возвращает хранимую там строку.
func (queue *Queue) PopWithNil() (string, bool) {
	if queue.first == nil {
		return "", false
	}
	item := queue.first
	queue.first = item.next
	item.next = nil
	if queue.first == nil {
		queue.last = nil
	}
	return item.value, true
}

// PopWithoutNil удаляет первый элемент из очереди и возвращает хранимую там строку.
func (queue *Queue) PopWithoutNil() (string, bool) {
	if queue.first == nil {
		return "", false
	}
	item := queue.first
	queue.first = item.next
	//item.next = nil
	if queue.first == nil {
		queue.last = nil
	}
	return item.value, true
}

// Push добавляет в конец очереди элемент с указанной строкой.
func (queue *Queue) Push(value string) {
	item := &QueueItem{value: value}
	if queue.first == nil { // нет элементов
		queue.first = item
		queue.last = item
		return
	}
	queue.last.next = item // текущий последний элемент должен указывать
	// на добавленный элемент
	queue.last = item // item становится последним элементом
}

func emulateQueueWithNil() {
	queue := &Queue{}

	queue.Push("0")
	queue.Push("1")
	queue.Push("2")
	for i := 3; i < 100_000_000; i++ {
		queue.PopWithNil()
		queue.Push(strconv.Itoa(i))
	}
}

func emulateQueueWithoutNil() {
	queue := &Queue{}

	queue.Push("0")
	queue.Push("1")
	queue.Push("2")
	for i := 3; i < 100_000_000; i++ {
		queue.PopWithoutNil()
		queue.Push(strconv.Itoa(i))
	}
}

func testMem(emulationFunc func()) {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	emulationFunc()
	runtime.ReadMemStats(&m2)
	funcName := runtime.FuncForPC(reflect.ValueOf(emulationFunc).Pointer()).Name()
	blue := color.New(color.FgHiBlue).SprintfFunc()
	fmt.Printf(
		"%s\n%s\n%s\n",
		blue(funcName),
		fmt.Sprintf("total bytes: %d", m2.TotalAlloc-m1.TotalAlloc),
		fmt.Sprintf("mallocs: %d", m2.Mallocs-m1.Mallocs),
	)
}

func RunTestMem() {
	testMem(emulateQueueWithNil)
	testMem(emulateQueueWithoutNil)
}
