package main

func brokenSearch(arr []int, value int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}

	leftIndex, rightIndex := 0, length - 1
	for rightIndex - leftIndex > 1 {
		midIndex := (rightIndex + leftIndex)/2
		right, mid, left := &arr[rightIndex], &arr[midIndex], &arr[leftIndex]

		if *left < *mid {
			if value >= *left && value <= *mid {
				rightIndex = midIndex
			} else {
				leftIndex = midIndex
			}
		} else if *mid < *right  {
			if value >= *mid && value <= *right {
				leftIndex = midIndex
			} else {
				rightIndex = midIndex
			}
		}
	}

	if arr[rightIndex] == value {
		return rightIndex
	}

	if arr[leftIndex] == value {
		return leftIndex
	}

	return -1
}
