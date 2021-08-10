package main

/*
type TreeInfo struct {
	diameter   int
	height     int
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeInfo.height + rightTreeInfo.height
	maxDiameterSoFar := max(leftTreeInfo.diameter, rightTreeInfo.diameter)
	currentDiameter := max(longestPathThroughRoot, maxDiameterSoFar)
	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)

	return TreeInfo{currentDiameter, currentHeight}
}
*/
