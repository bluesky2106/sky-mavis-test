package main

import (
	"fmt"
	"log"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoubleLinkedList struct {
	Len  int
	Tail *Node
	Head *Node
}

// Insert element to the end of the list
func (d *DoubleLinkedList) push_back(data int) {
	newNode := &Node{
		Data: data,
	}
	if d.Tail == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.Prev = d.Tail
		d.Tail.Next = newNode
		d.Tail = newNode
	}
	d.Len++
}

// Insert element to the beginning of the list
func (d *DoubleLinkedList) push_front(data int) {
	newNode := &Node{
		Data: data,
	}
	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.Next = d.Head
		d.Head.Prev = newNode
		d.Head = newNode
	}
	d.Len++
}

// get and remove the last element
func (d *DoubleLinkedList) pop_back() int {
	if d.Tail == nil {
		err := fmt.Errorf("The list is empty and there is no element to pop back")
		log.Fatal(err)
	}
	oldTail := d.Tail
	newTail := d.Tail.Prev
	if newTail == nil {
		d.Head = nil
		d.Tail = nil
	} else {
		newTail.Next = nil
		d.Tail = newTail
	}
	d.Len--
	return oldTail.Data
}

// get and remove the first element
func (d *DoubleLinkedList) pop_front() int {
	if d.Head == nil {
		err := fmt.Errorf("The list is empty and there is no element to pop front")
		log.Fatal(err)
	}
	oldHead := d.Head
	newHead := d.Head.Next
	if newHead == nil {
		d.Head = nil
		d.Tail = nil
	} else {
		newHead.Prev = nil
		d.Head = newHead
	}
	d.Len--
	return oldHead.Data
}

// traverse the list forward. used for checking the logic
func (d *DoubleLinkedList) traverse_forward() {
	if d.Head == nil {
		fmt.Println("Empty list")
		return
	}
	temp := d.Head
	for temp != nil {
		fmt.Printf("%d ", temp.Data)
		temp = temp.Next
	}
	fmt.Println()
}

func main() {
	dll := &DoubleLinkedList{}

	// raising error for the following 2 lines
	// dll.pop_back()
	// dll.pop_front()

	dll.push_back(1)
	dll.push_back(2)
	dll.push_back(3)
	dll.push_front(4)
	// now print the list to check if it is 4, 1, 2, 3?
	dll.traverse_forward()

	dll.pop_back()
	dll.pop_front()
	// now print the list to check if it is 1, 2?
	dll.traverse_forward()
}
