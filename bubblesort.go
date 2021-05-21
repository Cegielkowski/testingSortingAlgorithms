package main

import "time"

func BubbleSort(data []int) []int {
	defer TimeTrack(time.Now())

	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}




