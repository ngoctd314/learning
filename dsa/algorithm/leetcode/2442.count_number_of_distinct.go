package leetcode

func countDistinctIntegers(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	rev := func(x int) int {
		var revNum int
		for x > 0 {
			revNum = revNum*10 + x%10
			x = x / 10
		}
		return revNum
	}

	for i := range nums {
		m[nums[i]] = struct{}{}
		m[rev(nums[i])] = struct{}{}
	}

	return len(m)
}
