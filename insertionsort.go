package main

import "time"

func InsertionSort(data []int) []int {
	defer TimeTrack(time.Now())

	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}

	return data
}