package main

import (
	"errors"
	"fmt"
)

// DO NOT EDIT ---------------------------------------------------------------------------------------------------------

type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func printList(list *Node) {
	curNode := list
	for i := 0; curNode != nil; i++ {
		fmt.Printf("%d: %d\n", i, curNode.Value())
		curNode = curNode.Next()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// pushFront добавляет новый элемент в начало списка.
func pushFront(list *Node, value int) {
	secondItem := &Node{
		value: list.value,
		next:  list.next,
	}

	list.value = value
	list.next = secondItem
}

// pushBack добавляет новый элемент в конец списка.
func pushBack(list *Node, value int) {
	lastItem := &Node{
		value: value,
		next:  nil,
	}

	beforeLastItem := list
	for beforeLastItem.next != nil {
		beforeLastItem = beforeLastItem.next
	}
	beforeLastItem.next = lastItem
}

// count возвращает кол-во элементов в списке.
func count(list *Node) int {
	var counter int

	lastItem := list
	for lastItem != nil {
		lastItem = lastItem.next
		counter++
	}

	return counter
}

// popFront возвращает значение первого элемента и удаляет его из списка.
func popFront(list *Node) int {
	value := list.value

	list.value = list.next.value
	nextItem := list.next.next
	deletedItem := list.next
	list.next = nextItem
	deletedItem.next = nil

	return value
}

// popBack возвращает значение последнего элемента и удаляет его из списка.
func popBack(list *Node) int {
	beforeLastItem := list
	lastItem := list
	for lastItem.next != nil {
		beforeLastItem = lastItem
		lastItem = lastItem.next
	}

	value := lastItem.value
	lastItem = nil
	beforeLastItem.next = nil

	return value
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false.
func isValueInList(list *Node, value int) bool {
	var isValueFound bool

	item := list
	for item != nil {
		if item.value == value {
			isValueFound = true
			break
		}
		item = item.next
	}

	return isValueFound
}

// valueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range").
func valueByIndex(list *Node, index int) (int, error) {
	var counter, value int
	var err error

	if index > count(list)-1 {
		err = errors.New("index out of range")

		return value, err
	}

	item := list
	for item != nil {
		if counter == index {
			value = item.value
			break
		}
		item = item.next
		counter++
	}

	return value, err
}

// insert добавляет элемент в список в соответствующий индекс.
func insert(list *Node, index, value int) {
	if index > count(list)-1 {
		return
	}
	if index == count(list)-1 || index == count(list) {
		pushBack(list, value)
	}
	if index == 0 {
		pushFront(list, value)
	}

	var counter int

	beforeItem := list
	afterItem := list
	for afterItem != nil {
		beforeItem = afterItem
		afterItem = afterItem.next

		if index-counter == 1 {
			break
		}
		counter++
	}

	item := &Node{
		value: value,
		next:  afterItem,
	}
	beforeItem.next = item
}

// replace заменяет элемент в соответствующем индексе списка.
func replace(list *Node, index, value int) {
	if index > count(list)-1 {
		return
	}

	var counter int

	item := list
	for item != nil {
		if index == counter {
			break
		}
		item = item.next
		counter++
	}

	item.value = value
}

// deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range").
func deleteByIndex(list *Node, index int) (int, error) {
	var counter, value int
	var err error

	if index > count(list)-1 {
		err = errors.New("index out of range")

		return value, err
	}

	if index == count(list)-1 {
		pushBack(list, value)
	}
	if index == 0 {
		pushFront(list, value)
	}

	beforeItem := list
	item := list
	for item != nil {
		if counter == index {
			value = item.value
			beforeItem.next = item.next
			item = nil
			break
		}

		beforeItem = item
		item = item.next
		counter++
	}

	return value, err
}

// sort сортирует список (*).
func sort(list *Node) {
	length := count(list)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			leftValue, err := valueByIndex(list, j)
			if err != nil {
				return
			}
			rightValue, err := valueByIndex(list, j+1)
			if err != nil {
				return
			}

			if leftValue > rightValue {
				jValue, err := valueByIndex(list, j)
				if err != nil {
					return
				}
				j1Value, err := valueByIndex(list, j+1)
				if err != nil {
					return
				}
				replace(list, j, j1Value)
				replace(list, j+1, jValue)
			}
		}
	}
}

func main() {
	// var linkedList *Node
	linkedList := &Node{
		value: 0,
		next:  nil,
	}
	pushFront(linkedList, 10)
	pushFront(linkedList, 20)
	pushBack(linkedList, -10)
	pushBack(linkedList, -20)

	printList(linkedList)

	fmt.Println(count(linkedList))
	fmt.Println()

	fmt.Println(popFront(linkedList))
	printList(linkedList)

	fmt.Println(count(linkedList))
	fmt.Println()

	fmt.Println(popBack(linkedList))
	printList(linkedList)

	fmt.Println(count(linkedList))
	fmt.Println()

	fmt.Println(isValueInList(linkedList, 10))
	fmt.Println()

	fmt.Println(valueByIndex(linkedList, 0))
	fmt.Println()

	insert(linkedList, 1, 15)
	printList(linkedList)
	fmt.Println()

	fmt.Println(deleteByIndex(linkedList, 2))
	printList(linkedList)
	fmt.Println()

	sort(linkedList)
	printList(linkedList)
}
