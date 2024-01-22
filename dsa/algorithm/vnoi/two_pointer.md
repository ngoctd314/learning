# Two pointers

Kỹ thuật hai con trỏ được sử dụng khá phổ biến, giúp chương trình tiết kiệm thời gian và không gian xử lý.

## 1. Merge 2 mảng đã sắp xếp (không giảm) thành một mảng sắp xếp không giảm.

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

## 2. Two sum

Cho một mảng số nguyên a có n phần tử, mảng này đã được sắp xếp tăng dần. Hãy tìm hai vị trí khác nhau bất kỳ sao cho tổng của hai phần tử ở hai vị trí đó có giá trị là x.

```go
func twoSum(arr []int, acc int) [2]int {
	var rs [2]int

	i, j := 0, len(arr)-1
	for i < j {
		s := arr[i] + arr[j]
		if s == acc {
			rs[0], rs[1] = i, j
			return rs
		} else if s > acc {
			j--
		} else {
			i++
		}
	}

	return rs
}
```

## 3. Sub array

Cho dãy số nguyên dương a có n phần tử. Hãy tìm độ dài đoạn con dài nhất trong dãy sao cho tổng các phần tử trong đoạn này không quá s.


