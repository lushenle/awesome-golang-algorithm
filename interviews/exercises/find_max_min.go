package exercises

// FindMax Find the maximum element in a sorted and rotated list
func FindMax(nums []int) int {
	low, mid, high := 0, 0, len(nums)-1

	if nums[low] == nums[high] {
		return nums[low]
	}

	for low+1 < high {
		mid = low + (high-low)>>1
		if nums[low] < nums[mid] {
			low = mid
		} else if nums[low] > nums[mid] {
			high = mid
		} else {
			low = mid
		}
	}

	if nums[high] >= nums[low] {
		return nums[high]
	} else {
		return nums[low]
	}
}

// FindMin Find the minimum element in a sorted and rotated list
func FindMin(nums []int) int {
	low, mid, high := 0, 0, len(nums)-1

	for low+1 < high {
		mid = low + (high-low)>>1

		if nums[mid] == nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid
		} else {
			high = mid
		}
	}

	if nums[low] <= nums[high] {
		return nums[low]
	}

	return nums[high]
}
