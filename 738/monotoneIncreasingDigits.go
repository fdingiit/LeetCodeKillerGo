package monotoneIncreasingDigits

func doMonotoneIncreasingDigits(nums []int, begin, end int) {
	for i := begin; i < end; i++ {
		j := i + 1

		if nums[i] > nums[j] {
			for x := j; x <= end; x++ {
				nums[x] = 9
			}

			nums[i] -= 1
			doMonotoneIncreasingDigits(nums, 0, i)
		}
	}
}

func monotoneIncreasingDigits(N int) int {
	var nums []int
	var n int

	for N > 0 {
		nums = append(nums, N%10)
		N /= 10
	}

	// reverse
	for i := len(nums)/2 - 1; i >= 0; i-- {
		opp := len(nums) - 1 - i
		nums[i], nums[opp] = nums[opp], nums[i]
	}

	doMonotoneIncreasingDigits(nums, 0, len(nums)-1)

	for _, i := range nums {
		n = n*10 + i
	}

	return n
}
