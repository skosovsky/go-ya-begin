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
		v, ok := queue.Pop()
		if !ok {
			break
		}
		list = append(list, v)
	}
	fmt.Print(strings.Join(list, " "))
}

// Queue описывает очередь.
type Queue struct {
	front *QueueItem // указатель на первый элемент очереди
	rear  *QueueItem // указатель на последний элемент очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	value string     // данные
	next  *QueueItem // указатель на следующий элемент
}

// Pop удаляет первый элемент из очереди и возвращает хранимую там строку.
func (q *Queue) Pop() (string, bool) {
	if q.front == nil {
		return "", false
	}

	item := q.front
	value := item.value
	q.front = item.next
	item = nil
	if q.front == nil {
		q.rear = nil
	}

	return value, true
}

// Push добавляет в конец очереди элемент с указанной строкой.
func (q *Queue) Push(value string) {
	item := &QueueItem{value: value}
	if q.front == nil { // нет элементов
		// очередь пустая, поэтому добавляемый элемент
		// станет и первым и последним
		q.front = item
		q.rear = item
		return
	}
	q.rear.next = item // текущий последний элемент должен указывать
	// на добавленный элемент
	q.rear = item // item становится последним элементом
}
