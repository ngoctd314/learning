Full table là trường hợp mà optimizer phải đọc qua toàn bộ table để thực hiện câu query.

Dấu hiệu:

- Full index scan, full table scan.

**Trường hợp 1: rows sent lớn, filtered thấp (câu query lấy gần như là cả table)**

+ Có thể tính đến cách dùng limit offset để lấy về theo cách phân trang. Tuy nhiên cũng có nhiều hạn chế nếu không có trường id tự tăng.

**Trường hợp 2: rows sent thấp, filtered cao (câu query chỉ lấy một phần nhỏ)**
