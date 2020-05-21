package datastructures

import "fmt"

//SLNodeInt SL = Single Link
//SLNodeInt contains struct to store a Single Linked list struct containing a int Value
type SLNodeInt struct {
	Next  *SLNodeInt
	Value int
}

//SLMakeNode Makes new node
func SLMakeNode() *SLNodeInt {

	return new(SLNodeInt)
}

//GetValue Returns Value from SLNodeInt
func (node *SLNodeInt) GetValue() int {
	return node.Value
}

//SLAppendNode adds node to end of the list
func SLAppendNode(list *SLNodeInt, node *SLNodeInt) {
	for list.Next != nil {
		list = list.Next
	}
	list.Next = node
}

//SLDisplayValues display all values on a list
func SLDisplayValues(node *SLNodeInt) {
	for node != nil {
		fmt.Print(node.Value)
		node = node.Next
	}
}

//SLInsertFirst adds node to the start of a line
func SLInsertFirst(list *SLNodeInt, node *SLNodeInt) {
	node.Next = list
}

//SLFindPrevious finds previous node
func SLFindPrevious(head *SLNodeInt, node *SLNodeInt) *SLNodeInt {
	if head == nil || node == nil {
		return nil
	}
	for head.Next != node || head.Next != nil {
		head = head.Next
	}
	if head.Next == node {
		return head
	}
	return nil
}

//SLSwapNodes adds node to the list in the right spot
func SLSwapNodes(head *SLNodeInt, nodeA *SLNodeInt, nodeB *SLNodeInt) {
	for head != nil {
		if head == nodeA {

		}
		if head == nodeB {

		}
	}
}
