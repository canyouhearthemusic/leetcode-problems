package main

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1, nums2   []int
		expectedOutput float64
	}{
		{
			nums1:          []int{1, 3},
			nums2:          []int{2},
			expectedOutput: 2.0,
		},
		{
			nums1:          []int{1, 2},
			nums2:          []int{3, 4},
			expectedOutput: 2.5,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			actualOutput := findMedianSortedArrays(tt.nums1, tt.nums2)

			if actualOutput != tt.expectedOutput {
				t.Errorf("got %v, wanted %v", actualOutput, tt.expectedOutput)
			}
		})
	}
}
