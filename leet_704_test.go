package src

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	if search(arr, target) != 4 {
		t.Error("Binary search failed")
	}
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 2

	if search(arr, target) != -1 {
		t.Error("Binary search failed")
	}
}

func search(nums []int, target int) int {
	var start = 0
	var end = len(nums)

	for start < end {
		var middle = (start + end) / 2
		switch {
		case nums[middle] == target:
			{
				return middle
			}
		case middle == start || middle == end:
			{
				return -1
			}
		case nums[middle] > target:
			{
				end = middle
			}
		default:
			{
				start = middle
			}
		}
	}
	return -1
}
