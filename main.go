package main

func main() {
	myTree := BinaryTree{Value: 1}
	myTree.Left = &BinaryTree{Value: 2}
	myTree.Right = &BinaryTree{Value: 3}
	myTree.Right.Right = &BinaryTree{Value: 7}
	myTree.Right.Left = &BinaryTree{Value: 6}
	myTree.Left.Left = &BinaryTree{Value: 4}
	myTree.Left.Right = &BinaryTree{Value: 5}
	myTree.Left.Left.Left = &BinaryTree{Value: 8}
	myTree.Left.Left.Right = &BinaryTree{Value: 9}
	myTree.InvertBinaryTree()

}
