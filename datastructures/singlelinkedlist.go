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
	var aux = node
	for aux != nil {

		fmt.Print("->")
		fmt.Print(aux.Value)
		aux = aux.Next
	}
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
	if head == nil || nodeA == nil || nodeB == nil {
		return
	}
	var previousA = nodeA.SLFindPrevious(head)
	var previousB = nodeB.SLFindPrevious(head)
	var nextA = nodeA.Next
	var nextB = nodeB.Next

	previousB.Next = nodeA
	previousA.Next = nodeB

	nodeA.Next = nextB
	nodeB.Next = nextA
}

//SLbubbleSort sorts single linked list using bubble sort algorithm
func SLbubbleSort(head *SLNodeInt) {

}
