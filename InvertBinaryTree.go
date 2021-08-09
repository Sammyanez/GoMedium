package main

/*
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}
*/
func (tree *BinaryTree) InvertBinaryTree() {
	tree.invertbinaryTree()
}

func (tree *BinaryTree) invertbinaryTree() {
	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}
	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}

	if tree.Left == nil && tree.Right == nil {
		return
	} else {
		tree.Left, tree.Right = tree.Right, tree.Left
	}

}
