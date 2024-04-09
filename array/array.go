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

func singleNumber(nums []int) (n int) {
	for i := 0; i < len(nums); i++ {
		n ^= nums[i]
	}
	return n
}

func intersect(nums1 []int, nums2 []int) []int {
	var seen = make(map[int]int)
	for _, num := range nums1 {
		seen[num]++
	}

	var diff = make([]int, 0)
	for _, num := range nums2 {
		if n, load := seen[num]; load && n > 0 {
			diff = append(diff, num)
			seen[num]--
		}
	}
	return diff
}

func moveZeroes(nums []int) {
	z := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[z] = nums[i]
			z++
		}
	}
	for z < len(nums) {
		nums[z] = 0
		z++
	}
}
