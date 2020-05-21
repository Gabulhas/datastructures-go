package main

import (
	"datastrucures_go/datastructures"
)

func main() {
	head := datastructures.SLMakeNode()
	head.Value = 99
	for i := 0; i < 10; i++ {
		node := datastructures.SLMakeNode()
		node.Value = i
		datastructures.SLAppendNode(head, node)
	}
	datastructures.SLDisplayValues(head)
	//	datastructures.SLFindPrevious(head, node)
}
