package main

import "fmt"

func arrayMerger(arr1 [5]int, arr2 [4]int) [9]int {
	index1, index2 := 0, 0
	checkFirst, checkSecond := true, true
	var resultArr [9]int
	for i := 0; i < len(resultArr); i++ {
		if arr1[index1] <= arr2[index2] && checkFirst || !checkSecond {
			resultArr[i] = arr1[index1]

			if index1 == len(arr1)-1 {
				checkFirst = false
			}

			if index1 < len(arr1)-1 {
				index1++
			}

		} else {
			resultArr[i] = arr2[index2]
			if index2 == len(arr2)-1 {
				checkSecond = false
			}

			if index2 < len(arr2)-1 {
				index2++
			}

		}
	}

	return resultArr
}

func main() {
	firstArr := [5]int{1, 2, 5, 6, 7}
	secondArr := [4]int{2, 3, 4, 8}
	result := arrayMerger(firstArr, secondArr)
	fmt.Println(result)
}
