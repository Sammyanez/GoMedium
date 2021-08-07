package main

func (tree *BST) InOrderTraverse(array []int) []int {
	tree.inordertraverse(&array)
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	tree.preordertraverse(&array)
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	tree.postordertraverse(&array)
	return array
}

func (tree *BST) inordertraverse(array *[]int) {

	if tree.Left != nil {
		tree.Left.inordertraverse(array)
	}
	*array = append(*array, tree.Value)

	if tree.Right != nil {
		tree.Right.inordertraverse(array)
	}
}

func (tree *BST) preordertraverse(array *[]int) {

	*array = append(*array, tree.Value)
	if tree.Left != nil {
		tree.Left.preordertraverse(array)
	}

	if tree.Right != nil {
		tree.Right.preordertraverse(array)
	}
}

func (tree *BST) postordertraverse(array *[]int) {

	if tree.Left != nil {
		tree.Left.postordertraverse(array)
	}

	if tree.Right != nil {
		tree.Right.postordertraverse(array)
	}

	*array = append(*array, tree.Value)
}
