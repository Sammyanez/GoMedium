package main

func ReconstructBst(preOrderTraversalValues []int) *BST {
	myBST := BST{Value: preOrderTraversalValues[0]}
	for i := 1; i < len(preOrderTraversalValues); i++ {
		myBST.Insert(preOrderTraversalValues[i])
	}
	return &myBST
}
