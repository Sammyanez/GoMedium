package main

func main() {
	myTree := BinaryTree{Value: 1}
	myTree.Left = &BinaryTree{Value: 2}
	myTree.Right = &BinaryTree{Value: 3}
	myTree.Left.Left = &BinaryTree{Value: 4}
	myTree.Left.Left.Left = &BinaryTree{Value: 6}
	myTree.Left.Right = &BinaryTree{Value: 5}
	FindSuccessor(&myTree, myTree.Left.Right)
	print(BinaryTreeDiameter(&myTree))

}
