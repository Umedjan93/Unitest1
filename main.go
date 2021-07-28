package main

import "fmt"


func SliceAverage(Slice[]int)(int)  {
	SliceSum := 0
	AverageValue := 0
	for _, value := range Slice{
		SliceSum += value
		AverageValue = SliceSum/len(Slice)
	}
	return AverageValue
}

func main() {
	fmt.Println("Average slice value:",	SliceAverage([]int{1,2,3,4,5}))
}
