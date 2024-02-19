package main

import (
	"fmt"
	"strings"
)

func main() {
	list := []string{"На", "золотом", "крыльце", "сидели:", "царь,", "царевич,", "король,", "королевич."}
	queue := &Queue{}

	for _, v := range list {
		queue.Push(v)
	}
	list = list[:0]
	for {
		v, ok := queue.Pull()
		if !ok {
			break
		}
		list = append(list, v)
	}
	fmt.Print(strings.Join(list, " "))
}

// Queue описывает очередь.
type Queue struct {
	head   *QueueItem // указатель на головной элемент очереди
	tail   *QueueItem // указатель на хвостовой элемент очереди
	length int        // длина очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	value string     // данные
	next  *QueueItem // указатель на следующий элемент
}

// Pull удаляет первый элемент из очереди и возвращает хранимую там строку.
func (q *Queue) Pull() (string, bool) {
	if q.head == nil {
		return "", false
	}
	item := q.head
	value := item.value
	q.head = q.head.next
	item = nil
	q.length--

	if q.head == nil {
		q.tail = nil
	}

	return value, true
}

// Push добавляет в конец очереди элемент с указанной строкой.
func (q *Queue) Push(value string) {
	item := &QueueItem{
		value: value,
		next:  nil,
	}

	if q.head == nil { // очередь пустая, поэтому добавляемый элемент, станет и первым и последним
		q.head = item
		q.tail = item
		return
	}

	q.tail.next = item
	q.tail = item
	q.length++
}
