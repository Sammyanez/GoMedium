package main

func main() {
	myTree := BinaryTree{Value: 1}
	myTree.Left = &BinaryTree{Value: 2}
	myTree.Right = &BinaryTree{Value: 2}
	myTree.Left.Left = &BinaryTree{Value: 7}
	myTree.Left.Left.Left = &BinaryTree{Value: 8}
	myTree.Left.Left.Left.Left = &BinaryTree{Value: 9}
	myTree.Left.Right = &BinaryTree{Value: 4}
	myTree.Left.Right.Right = &BinaryTree{Value: 5}
	myTree.Left.Right.Right.Right = &BinaryTree{Value: 6}
	print(BinaryTreeDiameter(&myTree))

}
