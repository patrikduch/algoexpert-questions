package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) float64 {
	// Write your code here.

	closest := math.Inf(1)
	return findTheClosestValueInBstHelper(tree, target, closest)
}

func findTheClosestValueInBstHelper(tree *BST, target int, closest float64) float64 {

	if tree == nil { // We are on the leaf node
		return closest
	}

	if math.Abs(float64(target)-closest) > math.Abs(float64(target-tree.Value)) {

		closest = float64(tree.Value)
	}

	if target < tree.Value {

		return findTheClosestValueInBstHelper(tree.Left, target, closest)

	} else if target > tree.Value {

		return findTheClosestValueInBstHelper(tree.Right, target, closest)

	} else {
		return closest
	}
}

func main() {
	root := BST{Value: 6, Left: &BST{
		Value: 4,
		Right: &BST{
			Value: 8,
			Left:  nil,
			Right: nil,
		},
	}}

	result := root.FindClosestValue(2)

	fmt.Println(result)

}
