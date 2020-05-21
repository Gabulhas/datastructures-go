package main

import (
	"datastrucures_go/datastructures"
	"fmt"
)

func main() {
	head := datastructures.SLMakeNode()
	var temp *datastructures.SLNodeInt
	head.Value = 99
	for i := 0; i < 10; i++ {
		node := datastructures.SLMakeNode()
		node.Value = i
		datastructures.SLAppendNode(head, node)
		if i == 3 {
			temp = node
		}
	}
	datastructures.SLDisplayValues(head)
	var prev = datastructures.SLFindPrevious(head, temp)
	fmt.Printf("\n%d", prev.Value)
	datastructures.SLDisplayValues(head)

}
