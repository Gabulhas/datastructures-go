package datastructures

import (
	"errors"
	"fmt"
)

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
	var aux = node
	for aux != nil {

		fmt.Print("->")
		fmt.Print(aux.Value)
		aux = aux.Next
	}
	fmt.Printf("\n")
}

//SLInsertFirst adds node to the start of a line
func SLInsertFirst(list *SLNodeInt, node *SLNodeInt) {
	node.Next = list
}

//SLFindPrevious finds previous node
func (node *SLNodeInt) SLFindPrevious(head *SLNodeInt) *SLNodeInt {
	var aux = head
	if aux == nil || node == nil {
		return nil
	}
	for aux != nil {
		if aux.Next == node {
			return aux
		}
		aux = aux.Next
	}

	return nil
}

//SLSwapNodes adds node to the list in the right spot
func SLSwapNodes(head *SLNodeInt, nodeA *SLNodeInt, nodeB *SLNodeInt) {
	if head == nil || nodeA == nil || nodeB == nil || nodeA == nodeB {
		return
	}
	if nodeB == nodeA {
		return
	}
	nodeAPrev := nodeA.SLFindPrevious(head)
	nodeBPrev := nodeB.SLFindPrevious(head)
	if nodeAPrev != nil {
		nodeAPrev.Next = nodeB
	}
	if nodeBPrev != nil {
		nodeBPrev.Next = nodeA
	}
	tem := nodeA.Next
	nodeA.Next = nodeB.Next
	nodeB.Next = tem

}

//SLBubbleSort sorts single linked list using bubble sort algorithm
func SLBubbleSort(head *SLNodeInt) {
	var flag bool = true
	for flag {
		var aux = head
		var secondFlag bool = true
		for head != nil {
			fmt.Printf("%d", head.GetValue())
			fmt.Printf("%d", head.Next.GetValue())
			if head.Value > head.Next.Value {
				SLSwapNodes(aux, head, head.Next)
				secondFlag = false
			}
			head = head.Next
		}

		head = aux
		if secondFlag {
			flag = false
		}
		secondFlag = true
	}
}

//SLremoveNode removes node from given head/list
func (node *SLNodeInt) SLremoveNode(head *SLNodeInt) {
	if node == nil || head == nil {
		return
	}
	previous := node.SLFindPrevious(head)
	if previous != nil {

		previous.Next = node.Next
	}
	node = nil

}

//SLAddNext adds node in front of the current node
func (node *SLNodeInt) SLAddNext(NodeA *SLNodeInt) {
	if node == nil || NodeA == nil {
		return
	}
	temp := node.Next
	node.Next = NodeA
	NodeA.Next = temp

}

//SLFind return first node with a value
func SLFind(value int, head *SLNodeInt) (*SLNodeInt, error) {

	if head == nil {
		return nil, errors.New("invalid_head")
	}
	aux := head

	for aux != nil {

		if aux.GetValue() == value {
			return aux, nil
		}
		aux = aux.Next
	}

	return nil, errors.New("not_found")

}

//SLFindLoop finds if list is a loop using Floyd's Tortoise and Hare algorithm
func SLFindLoop(head *SLNodeInt) bool {

	if head == nil {
		return false
	}
	p1 := head
	p2 := head

	for p1.Next != nil && p2.Next.Next != nil {

		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			p1 = head
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next

			}
			return true
		}

	}
	return false
}
