package main

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	var myArray [][]int
	intervals = sort2D(intervals)
	for i := 0; i <= len(intervals)-1; i++ {
		if i == 0 {
			myArray = append(myArray, intervals[i])
		} else {
			if len(myArray) != 0 {
				myArray = Overlap(myArray[len(myArray)-1], intervals[i], myArray)
			}
		}
	}
	return myArray
}

func Overlap(array1 []int, array2 []int, array [][]int) [][]int {
	oneDArray := []int{}
	if array2[0] <= array1[1] && array1[1] <= array2[1] {
		if array1[0] < array2[0] {
			oneDArray = append(oneDArray, array1[0])
		} else {
			oneDArray = append(oneDArray, array2[0])
		}
		if array2[1] > array1[1] {
			oneDArray = append(oneDArray, array2[1])
		} else {
			oneDArray = append(oneDArray, array1[1])
		}

		array = array[:len(array)-1]
		array = append(array, oneDArray)
	} else if array1[0] <= array2[1] && array2[1] <= array1[1] {
		if array1[0] < array2[0] {
			oneDArray = append(oneDArray, array1[0])
		} else {
			oneDArray = append(oneDArray, array2[0])
		}
		if array2[1] > array1[1] {
			oneDArray = append(oneDArray, array2[1])
		} else {
			oneDArray = append(oneDArray, array1[1])
		}

		array = array[:len(array)-1]
		array = append(array, oneDArray)
	} else {
		array = append(array, array2)
	}
	return array
}

func sort2D(intervals [][]int) [][]int {

	for i := 0; i <= len(intervals)-1; i++ {
		smallestIndx := i
		smallestNum := intervals[i][0]
		for j := i; j <= len(intervals)-1; j++ {
			if i == j {
				continue
			}
			if smallestNum > intervals[j][0] {
				smallestIndx = j
				smallestNum = intervals[j][0]
			}
		}
		intervals[i], intervals[smallestIndx] = intervals[smallestIndx], intervals[i]
	}
	return intervals
}
