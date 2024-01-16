# VNOI Wiki

## Two pointers

Kỹ thuật hai con trỏ được sử dụng khá phổ biến, giúp chương trình tiết kiệm thời gian và không gian xử lý.

1. Merge 2 mảng đã sắp xếp (không giảm) thành một mảng sắp xếp không giảm.

```go
func merge2SortedArray(arr1, arr2 []int) []int {
	rs := make([]int, 0, len(arr1)+len(arr2))

	i, j, li, lj := 0, 0, len(arr1), len(arr2)
	for i < li || j < lj {
		if i == li || (j < lj && arr1[i] > arr2[j]) {
			rs = append(rs, arr2[j])
			j++
		} else {
			rs = append(rs, arr1[i])
			i++
		}
	}

	return rs
}
```

Độ phức tap O(n + m).
