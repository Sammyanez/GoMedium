package main

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	var numbersArray [][]int = [][]int{}
	sort.Ints(array)
	for i := range array{
		for j := i + 1; j < len(array)-1;j++{
			for x := j + 1; x <= len(array) - 1;x++{
				if array[i] + array[j] + array[x] == target{
					temp := []int{array[i],array[j],array[x]}
					numbersArray = append(numbersArray,temp)
				}
			}
		}
	}
	return numbersArray
}
