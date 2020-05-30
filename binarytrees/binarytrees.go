package datastructures

import (
	"errors"
	"fmt"
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