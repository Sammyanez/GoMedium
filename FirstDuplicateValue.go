package main

func FirstDuplicateValue(array []int) int {
	myArray := []int{}
	for i := range array {
		myArray = append(myArray, array[i])
		for j := range myArray {
			if i == j {
				continue
			}
			if array[i] == myArray[j] {
				return array[i]
			}
		}

	}
	return -1
}
