package leetcode

func peakIndexInMountainArray(arr []int) int {
	lo, hi := 1, len(arr)-2
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}
