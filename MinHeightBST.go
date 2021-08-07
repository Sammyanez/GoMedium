package main

func MinHeightBST(array []int) *BST {
	// Write your code here.
	return constructMinHeightBst(array, nil, 0, len(array)-1)
}

func constructMinHeightBst(array []int, bst *BST, startIdx int, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}
	midIdx := (startIdx + endIdx) / 2
	valueToAdd := array[midIdx]
	if bst == nil {
		bst = &BST{Value: valueToAdd}
	} else {
		bst.Insert(valueToAdd)
	}
	constructMinHeightBst(array, bst, startIdx, midIdx-1)
	constructMinHeightBst(array, bst, midIdx+1, endIdx)
	return bst
}
