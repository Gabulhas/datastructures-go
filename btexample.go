package main

import (
	dtbn "datastrucures_go/binarytrees"
	"fmt"
)

func main() {

	tree := dtbn.BTGenerateRandomTree(20, 4)
	fmt.Println(dtbn.BTSize(tree))
	dtbn.BTdisplayAllValues(tree)
	dtbn.BTdisplayAllValuesAndLevels(tree, 0)
}
