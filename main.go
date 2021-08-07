package main

func main() {
	myTree := BST{Value: 10}
	myTree.Insert(5)
	myTree.Insert(15)

	print(myTree.Contains(11))
}
