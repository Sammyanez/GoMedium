package main

func SpiralTraverse(array [][]int) []int {
	mySpiral := []int{}
	mMin := 0
	mMax := len(array) - 1
	nMin := 0
	nMax := len(array[0]) - 1
	iterateBack := false
	iteratedown := true
	goingAcross := true

	for mMin <= mMax && nMin <= nMax {
		if goingAcross {
			if iterateBack {
				for n := nMax; n >= nMin; n-- {
					mySpiral = appendability(mySpiral, array[mMax][n])
				}
				iterateBack = false
				mMax--
			} else {
				for n := nMin; n <= nMax; n++ {
					mySpiral = appendability(mySpiral, array[mMin][n])
				}
				iterateBack = true
				mMin++
			}
			goingAcross = false
		} else {
			if iteratedown {
				for m := mMin; m <= mMax; m++ {
					mySpiral = appendability(mySpiral, array[m][nMax])
				}
				iteratedown = false
				nMax--
			} else {
				for m := mMax; m >= mMin; m-- {
					mySpiral = appendability(mySpiral, array[m][nMin])
				}
				iteratedown = true

				nMin++
			}
			goingAcross = true
		}
	}
	return mySpiral
}

func appendability(array []int, i int) []int {
	if len(array) > 0 {
		if array[len(array)-1] == i {
			return array
		} else {
			return append(array, i)
		}
	} else {
		return append(array, i)
	}
}
