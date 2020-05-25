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
	//var prev = temp.SLFindPrevious(head)
	//fmt.Println(prev.Value)

	fmt.Print("\n")
	datastructures.SLSwapNodes(head, temp, tempS)
	datastructures.SLDisplayValues(head)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	datastructures.SLSwapNodes(head, tempS, head)
	fmt.Print("\n")

	datastructures.SLDisplayValues(head)
}
