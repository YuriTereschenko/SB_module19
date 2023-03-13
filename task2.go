package main

import "fmt"

func bubbleSort(arrToSort [6]int) [6]int {
	for rightShift := 0; rightShift < len(arrToSort); rightShift++ {
		for i := 1; i < len(arrToSort)-rightShift; i++ {
			if arrToSort[i] < arrToSort[i-1] {
				arrToSort[i], arrToSort[i-1] = arrToSort[i-1], arrToSort[i]
			}
		}
	}
	return arrToSort
}

func main() {
	arr := [6]int{5, 2, 6, 4, 8, 3}
	res := bubbleSort(arr)
	fmt.Println(res)

}
