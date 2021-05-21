package main

import (
	"fmt"
	"time"
)

func main()  {
	unorderedTen := generateData(10, false)
	unorderedOneHundred := generateData(100000, false)

	orderedTen := generateData(10, true)
	orderedOneHundred := generateData(100000, true)

	fmt.Println("=================== Block of 10 numbers unordered ===================")
	BubbleSort(unorderedTen)
	HeapSort(unorderedTen)
	InsertionSort(unorderedTen)
	mergeSort(unorderedTen)
	QuickSort(unorderedTen)


	fmt.Println("=================== Block of 100000 numbers unordered ===================")
	BubbleSort(unorderedOneHundred)
	HeapSort(unorderedOneHundred)
	InsertionSort(unorderedOneHundred)
	mergeSort(unorderedOneHundred)
	QuickSort(unorderedOneHundred)

	fmt.Println("=================== Block of 10 numbers ordered ===================")
	BubbleSort(orderedTen)
	HeapSort(orderedTen)
	InsertionSort(orderedTen)
	mergeSort(orderedTen)
	QuickSort(orderedTen)

	fmt.Println("=================== Block of 100000 numbers ordered ===================")
	BubbleSort(orderedOneHundred)
	HeapSort(orderedOneHundred)
	InsertionSort(orderedOneHundred)
	mergeSort(orderedOneHundred)
	QuickSort(orderedOneHundred)
}

func mergeSort(data []int) []int {
	defer TimeTrack(time.Now())
	return MergeSort(data)
}

func generateData(limit int, ordered bool) []int {
	var result []int

	for i := 1; i <= limit; i++ {
		if ordered {
			result = append(result, i)
		} else {
			result = append(result, Random())
		}
	}
	return result
}