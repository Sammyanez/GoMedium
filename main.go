package main

func main() {
	myTree := BST{Value: 10}
	myTree.Left = &BST{Value: 5}
	myTree.Right = &BST{Value: 15}
	myTree.Right.Right = &BST{Value: 22}
	myTree.Right.Left = &BST{Value: 13}
	myTree.Right.Left.Right = &BST{Value: 14}
	myTree.Left.Left = &BST{Value: 2}
	myTree.Left.Right = &BST{Value: 5}
	myTree.Left.Left.Left = &BST{Value: 1}

	print(FindKthLargestValueInBst(&myTree, 2))
}
