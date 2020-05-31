package main

import (
	dtsl "datastrucures_go/singlelinkedlist"
)

func main() {
	head := dtsl.SLMakeNode()
	var temp *dtsl.SLNodeInt
	var tempS *dtsl.SLNodeInt

	head.Value = 99
	for i := 0; i < 10; i++ {
		node := dtsl.SLMakeNode()
		node.Value = i
		dtsl.SLAppendNode(head, node)
		if i == 3 {
			temp = node
		}
		if i == 7 {
			tempS = node
		}
	}
	dtsl.SLDisplayValues(head)
	temp.SLremoveNode(head)
	dtsl.SLDisplayValues(head)
	tempS.SLremoveNode(head)
	dtsl.SLDisplayValues(head)
	aux := head
	for aux != nil {
		nextNode := aux.Next
		if aux.GetValue()%2 == 0 {
			aux.SLremoveNode(head)
		}
		aux = nextNode
	}
	dtsl.SLDisplayValues(head)
	tempF, _ := dtsl.SLFind(5, head)
	tempH := dtsl.SLMakeNode()
	tempH.Value = 192
	tempF.SLAddNext(tempH)

	dtsl.SLDisplayValues(head)

}
