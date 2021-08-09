package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree // is previous
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	var myTrees []*BinaryTree
	findsuccessor(tree, node, &myTrees)
	for i := range myTrees {
		if myTrees[i] == node {
			if i >= len(myTrees)-1 {
				break
			}
			return myTrees[i+1]
		}
	}
	return nil
}

func findsuccessor(tree *BinaryTree, node *BinaryTree, trees *[]*BinaryTree) {

	if tree.Left != nil {
		findsuccessor(tree.Left, node, trees)
	}

	*trees = append(*trees, tree)

	if tree.Right != nil {
		findsuccessor(tree.Right, node, trees)
	}

}
