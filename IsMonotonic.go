package main

func IsMonotonic(array []int) bool {
	decreasing, increasing := false, false
	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			increasing = true
		} else if array[i] > array[i+1] {
			decreasing = true
		} else {
			continue
		}
		if increasing && decreasing == false {
			continue
		} else if decreasing && increasing == false {
			continue
		} else {
			return false
		}
	}
	return true
}
