package main

func LongestPeak(array []int) int {
	peakLength := 0
	first := true
	neverDecreasing := true
	temp := 1
	for i := 0; i <= len(array)-1; i++ {
		first = true
		for i != len(array)-1 && array[i] < array[i+1] {
			first = false
			temp++
			i++
		}
		for !first && i < len(array)-1 && array[i] > array[i+1] {
			neverDecreasing = false
			temp++
			i++
		}
		if peakLength < temp {
			peakLength = temp
			temp = 1
		} else {
			if temp > 1 {
				i--
			}
			temp = 1
		}
	}

	if neverDecreasing {
		return 0
	}
	return peakLength
}
