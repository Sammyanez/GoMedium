package main

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	var smallest int
	var smallestNumbers []int
	for i := range array1 {
		for j := 0; j < len(array2); j++ {
			if i == 0 && j == 0 {
				smallest = int(math.Abs(float64(array1[i] - array2[j])))
				smallestNumbers = append(smallestNumbers, array1[i])
				smallestNumbers = append(smallestNumbers, array2[j])
			} else {
				if smallest > int(math.Abs(float64(array1[i]-array2[j]))) {
					smallest = int(math.Abs(float64(array1[i] - array2[j])))
					smallestNumbers[0] = array1[i]
					smallestNumbers[1] = array2[j]
				}
			}

		}
	}
	return smallestNumbers
}
