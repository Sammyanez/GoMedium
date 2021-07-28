package main

import "sort"

func MoveElementToEnd(array []int, toMove int) []int {
	notOurNum := []int{}
	OurNum := []int{}
	sort.Ints(array)
	for i := range array {
		if array[i] == toMove {
			OurNum = append(OurNum, array[i])
		} else {
			notOurNum = append(notOurNum, array[i])
		}
	}
	return append(notOurNum, OurNum...)
}
