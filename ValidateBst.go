package main

import "math"

func (tree *BST) ValidateBst() bool {
	return tree.validateBst(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validateBst(min int, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.validateBst(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validateBst(tree.Value, max) {
		return false
	}
	return true
}

/* does not check every level so it doesnt work

func (tree *BST) ValidateBst() bool {
	if tree.Left.validateLeft(tree.Value)&& tree.Right.validateRight(tree.Value) {
		return true
	}
	return false
}

func (tree *BST)validateLeft( rootValue int) bool {
	if tree == nil{
		return true
	}
	if rootValue <= tree.Value{
		return false
	}
	if tree.Left!=nil{
		if tree.Left.Value >= tree.Value{
			return false
		}
	}
	if tree.Right != nil{
		if tree.Value > tree.Right.Value{
			return false
		}
	}

	return tree.Left.validateLeft(rootValue) && tree.Right.validateLeft(rootValue)

}

func (tree *BST)validateRight( rootValue int) bool {
	if tree == nil{
		return true
	}
	if rootValue > tree.Value{
		return false
	}
	if tree.Left!=nil{
		if tree.Left.Value >= tree.Value{
			return false
		}
	}
	if tree.Right != nil{
		if tree.Value > tree.Right.Value{
			return false
		}
	}

	return tree.Left.validateRight(rootValue) && tree.Right.validateRight(rootValue)
}

*/
