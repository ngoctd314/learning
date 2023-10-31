package leetcode

func findKthPositive(arr []int, k int) int {
	k -= arr[0] - 1
	search := func(lo, hi int, target int) int {
		for lo <= hi {
			mid := (lo + hi) / 2
			if mid == target {
				return mid
			} else if mid > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return -1
	}
	if k <= 0 {
		return k + arr[0] - 1
	}

	for i := 1; i <= len(arr)-1; i++ {
		tmp := arr[i] - arr[i-1] - 1
		if tmp > 0 {
			k -= tmp
			if k <= 0 {
				return search(arr[i-1]+1, arr[i]-1, k+tmp+arr[i-1])
			}
		}
	}

	return arr[len(arr)-1] + k
}
