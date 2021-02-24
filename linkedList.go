package main

import (
	"errors"
	"fmt"
)

/* Linked List -
1. To Create a Node struct which contains value and the pointer to the next node.
2. To Create a linked list struct that contains a pointer reference to Node struct and the length of the
   linked list.
*/

// Node Structure
type Node struct {
	value interface{}
	next  *Node
}

// LinkedList Structure
type LinkedList struct {
	head   *Node
	length int
}

/* Linked List Operations -
1. Insert a value into linked list.
2. Traversing the linked list.
3. Delete the value from the linked list
*/

// Insert a value in the linked list.
func (l *LinkedList) Insert(value interface{}) {
	newNode := Node{}
	newNode.value = value
	if l.length == 0 {
		l.head = &newNode
		l.length++
		return
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			l.length++
			return
		}
		ptr = ptr.next
	}

}

// PrintList prints the node in the linked list.
func (l *LinkedList) PrintList() {
	if l.length == 0 {
		fmt.Println("List is Empty")
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println("Node :", ptr)
		ptr = ptr.next
	}
}

// GetAt the value at the perticular position in the linked list.
func (l *LinkedList) GetAt(pos int) *Node {
	if pos < 0 {
		fmt.Println("position cannot be negitive")
	}
	if pos > l.length-1 {
		return nil
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if i == pos {
			fmt.Printf("Node at the position %v : %v \n", pos, ptr)
			return ptr
		}
		ptr = ptr.next
	}
	return nil
}

// Search for the value in the list and return the position
func (l *LinkedList) Search(value interface{}) int {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.value == value {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// InsertAt insert's a node at a perticular position
func (l *LinkedList) InsertAt(pos int, value interface{}) {
	newNode := Node{}
	newNode.value = value
	if pos > l.length-1 {
		return
	}
	if pos < 0 {
		fmt.Println("Position cannot be negitive")
		return
	}

	n := l.GetAt(pos)
	newNode.next = n
	preNode := l.GetAt(pos - 1)
	preNode.next = &newNode
	l.length++
}

// DeleteAt deletes a Node from the list based on the position provides
func (l *LinkedList) DeleteAt(pos int) error {
	if pos > l.length-1 {
		return nil
	}
	if pos < 0 {
		fmt.Println("Position cannot be negitive")
		return nil
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.length--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.length == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.length; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.length--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

// func main() {
// 	list := LinkedList{}
// 	list.Insert(5)
// 	list.Insert(6)
// 	list.Insert(8)
// 	list.InsertAt(2, 7)
// 	fmt.Print("------BEFORE DELETE-------- \n")
// 	list.PrintList()
// 	// list.DeleteAt(3)
// 	list.DeleteVal(6)
// 	fmt.Print("------AFTER DELETE-------- \n")
// 	list.PrintList()
// 	// list.GetAt(0)
// 	// fmt.Printf("Position of the searched Node: %v \n", list.Search(7))
// }
