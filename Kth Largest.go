package main

import "sort"

func FindKthLargestValueInBst(tree *BST, k int) int {
	var myArray []int
	myArray = tree.InOrderTraverse(myArray)
	sort.Sort(sort.Reverse(sort.IntSlice(myArray)))

	return myArray[k]
}
