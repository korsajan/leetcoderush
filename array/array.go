package array

func removeDuplicates(nums []int) int {
	var i, n = 0, len(nums)
	for i < n {
		j := binarySearch(nums, i+1, n-1, nums[i])
		if j == n {
			break
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}

func binarySearch(a []int, low, high, target int) int {
	r := -1
	for low <= high {
		mid := (high + low) >> 1
		if a[mid] <= target {
			low = mid + 1
		} else {
			r = mid
			high = mid - 1
		}
	}
	if r == -1 {
		return len(a)
	}
	return r
}

func rotate(nums []int, k int) {
	n := len(nums)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[(i+k)%n] = nums[i]
	}
	copy(nums, out)
}

func containsDuplicate(nums []int) bool {
	var filter = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := filter[n]; ok {
			return ok
		}
		filter[n] = struct{}{}
	}
	return false
}
