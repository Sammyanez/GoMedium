package main

// This is an input class. Do not edit.
type TreeInfo struct {
	diameter   int
	height     int
	isBalanced bool
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	treeInfo := getTreeInfo(tree)
	return treeInfo.isBalanced
}

func getTreeInfo(node *BinaryTree) TreeInfo {
	if node == nil {
		return TreeInfo{isBalanced: true, height: -1}
	}

	leftSubtreeInfo := getTreeInfo(node.Left)
	rightSubtreeInfo := getTreeInfo(node.Right)

	isBalanced := leftSubtreeInfo.isBalanced && rightSubtreeInfo.isBalanced && abs(
		leftSubtreeInfo.height-rightSubtreeInfo.height) <= 1
	height := max(leftSubtreeInfo.height, rightSubtreeInfo.height) + 1

	return TreeInfo{
		isBalanced: isBalanced,
		height:     height,
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b

}
