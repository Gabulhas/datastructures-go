package main

import (
	"datastrucures_go/datastructures"
	"fmt"
)

func main() {
	head := datastructures.SLMakeNode()
	var temp *datastructures.SLNodeInt
	var tempS *datastructures.SLNodeInt

	head.Value = 99
	for i := 0; i < 10; i++ {
		node := datastructures.SLMakeNode()
		node.Value = i
		datastructures.SLAppendNode(head, node)
		if i == 3 {
			temp = node
		}
		if i == 7 {
			tempS = node
		}
	}
	datastructures.SLDisplayValues(head)
	temp.SLremoveNode(head)
	datastructures.SLDisplayValues(head)
	tempS.SLremoveNode(head)
	datastructures.SLDisplayValues(head)
	aux := head
	for aux != nil {
		nextNode := aux.Next
		if aux.GetValue()%2 == 0 {
			aux.SLremoveNode(head)
		}
		aux = nextNode
	}
	datastructures.SLDisplayValues(head)
	tempF, _ := datastructures.SLFind(5, head)
	tempH := datastructures.SLMakeNode()
	tempH.Value = 192
	tempF.SLAddNext(tempH)

	datastructures.SLDisplayValues(head)

	fmt.Println(datastructures.SLFindLoop(head))
}
