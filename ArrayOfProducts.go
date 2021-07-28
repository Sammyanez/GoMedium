package main

func ArrayOfProducts(array []int) []int {
	myArray := []int{}
	for i := range array {
		temp := 1
		for j := range array {
			if i == j {
				continue
			} else {
				temp *= array[j]
			}
		}
		myArray = append(myArray, temp)
	}
	return myArray
}
