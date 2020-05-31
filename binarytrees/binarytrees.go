package datastructures

import (
	"fmt"
	"math/rand"
)

//BTNodeInt SL = Single Link
//BTNodeInt contains struct to store a Single Linked list struct containing a int Value
type BTNodeInt struct {
	NextR *BTNodeInt
	NextL *BTNodeInt
	Value int
}

//BTMakeNode makes new
func BTMakeNode() *BTNodeInt {

	return new(BTNodeInt)

}

//BTGenerateRandomTree generates random tree and returns root given a max and min number of childNodes, made for testing the functions implemented in this packages
func BTGenerateRandomTree(maxValue int, maxLength int) *BTNodeInt {

	head := BTMakeNode()
	head.Value = randomInt(maxValue, 0)
	generate(head, maxValue, maxLength)

	return head
}
func generate(head *BTNodeInt, maxValue, maxLength int) {
	if maxLength == 0 {
		return
	}

	nodeR := BTMakeNode()
	nodeR.Value = randomInt(maxValue, 0)

	nodeL := BTMakeNode()
	nodeL.Value = randomInt(maxValue, 0)
	head.NextR = nodeR
	head.NextL = nodeL
	generate(nodeR, maxValue, randomInt(maxLength, 0))
	generate(nodeL, maxValue, randomInt(maxLength, 0))
}

//BTSize Returns size of tree
func BTSize(head *BTNodeInt) int {

	if head == nil {
		return 0
	}
	return 1 + BTSize(head.NextL) + BTSize(head.NextR)
}

//BTdisplayAllValues displays all values in a line. Just for testing purposes.
func BTdisplayAllValues(head *BTNodeInt) {
	if head == nil {
		return
	}

	fmt.Printf("%d ", head.Value)
	BTdisplayAllValues(head.NextR)
	BTdisplayAllValues(head.NextL)
}

//BTdisplayAllValuesAndLevels does what the name of the function says :D
func BTdisplayAllValuesAndLevels(head *BTNodeInt, level int) {
	if head == nil {
		return
	}

	fmt.Println(level, " ", head.Value)
	BTdisplayAllValuesAndLevels(head.NextR, level+1)
	BTdisplayAllValuesAndLevels(head.NextL, level+1)

}

func randomInt(max, min int) int {

	return rand.Intn(max-min+1) + min
}
