package main

import (
	"fmt"
	"math"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns kth node of linked list
func getKthNode(head *node, k int) *node {
	current := head

	for k > 0 && current != nil {
		current = current.next
		k--
	}

	return current
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Finds intersection node of two linked lists
func (list1 *linkedList) findIntersection(list2 *linkedList) *node {
	if list1 == nil || list2 == nil {
		return nil
	}

	tailList1 := list1.tail
	tailList2 := list2.tail

	lengthList1 := list1.length
	lengthList2 := list2.length

	if tailList1 != tailList2 {
		return nil
	}

	var shorter *node
	var longer *node

	if lengthList1 < lengthList2 {
		shorter = list1.head
	} else {
		shorter = list2.head
	}

	if lengthList1 < lengthList2 {
		longer = list2.head
	} else {
		longer = list1.head
	}

	longer = getKthNode(longer, int(math.Abs(float64(lengthList1)-float64(lengthList2))))

	for shorter != longer {
		shorter = shorter.next
		longer = longer.next
	}

	return longer
}

func main() {
	node1 := &node{data: 3}
	node2 := &node{data: 1}
	node3 := &node{data: 5}
	node4 := &node{data: 9}
	node5 := &node{data: 7}
	node6 := &node{data: 2}
	node7 := &node{data: 1}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)
	linkedList1.pushBack(node5)
	linkedList1.pushBack(node6)
	linkedList1.pushBack(node7)

	node8 := &node{data: 4}
	node9 := &node{data: 6}
	linkedList2 := linkedList{}

	linkedList2.pushBack(node8)
	linkedList2.pushBack(node9)
	linkedList2.pushBack(node5)
	linkedList2.pushBack(node6)
	linkedList2.pushBack(node7)

	result := linkedList1.findIntersection(&linkedList2)

	if result != nil {
		fmt.Println("The intersection value is: ", result.data)
	} else {
		fmt.Println("There's no intersection")
	}
}
