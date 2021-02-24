package main

import (
	"errors"
	"fmt"
)

/* Doubly Linked List -
1. To Create a Node struct which contains value, the pointer to the next node and the pointer to the prevous node.
2. To Create a linked list struct that contains a pointer reference to Node struct and the length of the
   linked list.
*/

// DLLNode Structure
type DLLNode struct {
	previous *DLLNode
	value    interface{}
	next     *DLLNode
}

// DoublyLinkedList Structure
type DoublyLinkedList struct {
	head   *DLLNode
	length int
}

// InsertDLL into DLL
func (dll *DoublyLinkedList) InsertDLL(value interface{}) {
	newNode := DLLNode{}
	newNode.value = value
	if dll.length == 0 {
		dll.head = &newNode
		dll.length++
		return
	}
	ptr := dll.head
	for i := 0; i < dll.length; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			newNode.previous = ptr
			dll.length++
			return
		}
		ptr = ptr.next
	}
}

// PrintDLList prints the node in the linked list.
func (dll *DoublyLinkedList) PrintDLList() {
	if dll.length == 0 {
		fmt.Println("List is Empty")
	}
	ptr := dll.head
	for i := 0; i < dll.length; i++ {
		fmt.Println("Node :", ptr)
		ptr = ptr.next
	}
}

// GetAt the value at the perticular position in the linked list.
func (dll *DoublyLinkedList) GetAt(pos int) *DLLNode {
	if pos < 0 {
		fmt.Println("position cannot be negitive")
	}
	if pos > dll.length-1 {
		return nil
	}
	ptr := dll.head
	for i := 0; i < dll.length; i++ {
		if i == pos {
			fmt.Printf("Node at the position %v : %v \n", pos, ptr)
			return ptr
		}
		ptr = ptr.next
	}
	return nil
}

// Search for the value in the list and return the position
func (dll *DoublyLinkedList) Search(value interface{}) int {
	ptr := dll.head
	for i := 0; i < dll.length; i++ {
		if ptr.value == value {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// InsertDLLAt insert Node into DLL in a perticular position
func (dll *DoublyLinkedList) InsertDLLAt(pos int, value interface{}) {
	if pos < 0 {
		fmt.Println("position cannot be -ve.")
		return
	}
	if pos > dll.length-1 {
		fmt.Println("position cannot be greater than the list itself.")
		return
	}

	newNode := DLLNode{}
	newNode.value = value

	n := dll.GetAt(pos)
	n.previous = &newNode
	prevNode := dll.GetAt(pos - 1)
	prevNode.next = &newNode
	newNode.previous = prevNode
	newNode.next = n
	dll.length++
}

// DeleteAt deletes a Node from the list based on the position provides
func (dll *DoublyLinkedList) DeleteAt(pos int) error {
	if pos > dll.length-1 {
		return nil
	}
	if pos < 0 {
		fmt.Println("Position cannot be negitive")
		return nil
	}
	prevNode := dll.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = dll.GetAt(pos).next
	dll.GetAt(pos).previous = prevNode.previous
	dll.length--
	return nil
}

// func main() {
// 	// list := DoublyLinkedList{}
// 	// list.InsertDLL(5)
// 	// list.InsertDLL(6)
// 	// list.InsertDLL(8)
// 	// list.InsertDLLAt(2, 7)

// 	// fmt.Println("-------BEFORE---------")
// 	// list.PrintDLList()
// 	// fmt.Println("----------------------")

// 	// list.DeleteAt(2)
// 	// fmt.Println("-------AFTER---------")
// 	// list.PrintDLList()
// 	// fmt.Println("----------------------")

// 	// list.GetAt(1)
// }
