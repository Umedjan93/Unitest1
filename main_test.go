package main

import "testing"

func TestSliceAverage(t *testing.T) {

	tests := []struct {
		name string
		Slice[]int
		AverageWant int
	}{
		{ name: "All of numbers are positive", Slice: []int{5,4,3,2,1}, AverageWant: 3},
		{ name: "All of numbers are negative", Slice: []int{-1,-2,-3,-4,-5}, AverageWant: -3},
		{name: "All numbers are mixed", Slice: []int{-1,1,-2,2,0},  AverageWant: 0},
	}
	for _, tt := range tests {
		AverageValue := SliceAverage(tt.Slice)
		{
			if AverageValue != tt.AverageWant {
				t.Errorf("Index of SliceAverage test %s, AverageValue %d and AverageWant %d",
					tt.name, AverageValue, tt.AverageWant)
			}
		}
	}
}