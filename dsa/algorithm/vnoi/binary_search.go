package vnoi

func binarySearch(ar []int, target int) int {
	l, h := 0, len(ar)-1
	for l < h {
		m := (l + h) / 2
		if ar[m] == target {
			return m
		}
		if ar[m] > target {
			h = m
		} else {
			l = m
		}
	}
	return -1
}

/*
Cơ sở lý thuyết: Định lý chính của tìm kiếm nhị phân
Khi một bài toán mà ta đoán được có thể dùng tìm kiếm nhị phân để giải, thì ta phải chứng minh tính đúng đắn suy luận của chúng ta. Do đó, xây dụng một cơ sở lý thuyết vững chắc là vô cùng cần thiết.

Cho không gian tìm kiếm S bao gồm các ứng cử viên cho kết quả của bài toán. Ta định nghĩa một hàm kiểm tra P: S -> true, false là hàm nhận một ứng cử viên x thuộc S và trả về giá trị true/false cho biết x có hợp lệ hay không. Hàm P là hàm kiểm tra một tính chất nào đó, xem một ứng cử viên cho kết quả của bài toán có thỏa tính chất đó không.

Một bài toán chỉ có thể áp dụng tìm kiếm nhị phân nếu và chỉ nếu hàm kiểm tra P của bài toán thỏa mãn
- Với mỗi x, y thuộc S. y > x và P(x) = true thì P(y) = true

Lưu ý rằng chính chất trên của hàm kiểm tra P cũng tương đương với tính chất
- Với mỗi x, y thuộc S. y < x và P(x) = false => P(y) = false

Từ định lý trên, ta rút ra được mấu chốt để giải một bài toán tìm kiếm nhị phân là ta cần thiết kế được hàm P hợp lý sao cho thỏa mãn điều kiện trong định lý chính.
*/

// https://vnoi.info/wiki/algo/basic/binary-search.md
func P(ar []int) func(target int) bool {
	return func(target int) bool {
		return false
	}
}

func binarySearchGeneral(ar []int, target int) {

}
