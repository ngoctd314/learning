# Cache

## Cache dưới góc nhìn của anh Son Luong N.

### 1. Câu chuyện của tác giả

Cache nó là 1 giải pháp cho vấn đề truy cập data trong hệ thống. Thường thì những vấn đề truy cập này có nhiều cách giải quyết khác nhau, tùy vào các thông số và yêu cầu của doanh nghiệp. Những vấn đề này thường có những giải pháp khác nhau, đa phần là có các điểm mạnh yếu riêng.

Ở cty cũ của mình, hệ thống và code chạy được hơn 20 năm ở một 1 e-commerce load (đọc nhiều, mua/write ít) nên cũng có thời dùng cache (memcached). Nhưng sau nhiều năm quan sát thì nhận ra là... Engineer nhìn chung khá dễ mắc lỗi trong lúc code khi có nhiều nguồn cho data, nguồn này khác nguồn kia. Và không có cách nào để giải quyết những lỗi như thế 1 cách triệt để. Engineer có thể đi cty khác, tuyển người mới vào -> lặp lại vấn đề. Sau này các bạn bên DB nhận ra là có thể tăng tốc DB bằng cách chỉnh cho 1 phần nóng của database chạy từ RAM, perf không thua cache. System thì đã có data sharding + main/replica sẵn rồi nên scale lên không khó. Hủy tầng cache làm hệ thống đơn giản hơn, giúp cho engineer đỡ mắc lỗi hơn.

### 2. Bài học của tác giả

Bài học ở đây là sự "đơn giản" cũng là 1 tính chất đáng quý của 1 hệ thống. Khi design thì nên chú ý tới để giảm rủi ro cho doanh nghiệp lâu dài.

Ở VN rất nhiều cty chạy theo những hệ thống mới, phức tạp để... Lên CV cho tech lead. Nhiều khi bỏ qua những giải pháp đơn giản cho dù mạnh yếu ngang nhau. Doanh nghiệp phải thay engineer mỗi năm vì thị trường tăng trưởng mạnh, nên rủi ro khá cao khi dùng hệ thống phức tạp.

Thế nên cái blog [này](https://blog.koehntopp.info/2021/03/12/memory-saturated-mysql.html) là 1 dẫn chứng khá hay: nếu 1 doanh nghiệp e-commerce trăm tỷ đô có thể chạy không cần cache, thì chưa chắc những doanh nghiệp nhỏ hơn đã cần.

Nhưng chém cho vui thế thôi, các bạn đừng vứt cục cache đi nhé. Mỗi doanh nghiệp mỗi khác, chỉ có thể copy cách suy luận chứ đừng copy giải pháp.

### 3. Phản biện

#### 4. Phản biện của anh Viet Tran

Cach là khái niệm chung của computer science không riêng cho system hay backend, db... Vì thế engineer làm ở mảng nào thì có trách nhiệm xử lý ở mảng đó. DB engine thì sẽ có cache problem của riêng nó, application cũng thế, từ backend tới client. Nên ông engineer DB không thể nói rằng DB đã có cache rồi thì các bạn không cần cache ở những tầng khác nữa. Vì dữ liệu cache là đặc thù của nghiệp vụ, bản thân cái DB chưa phải là dữ liệu cuối cùng.

Nhưng năm gần đây, engineer có xu hướng xem storage (DB, File System) chỉ nên làm nhiệm vụ lưu trữ của nó, không hơn không kém. Các engineer sẽ giảm lệ thuộc vào storage nhiều hơn.

Cache như thế nào cho đúng khác với có nên giảm cache hay không.

Bản thân e đi hỗ trợ doanh nghiệp nhiều nơi thì thấy rằng: các dev (VN rất ít engineer) không biết, không rõ nên làm thế nào nên phải tham khác các giải pháp từ công ty lớn. Có những giải pháp đơn giản nhưng đa số đều phức tạp với scale hiện có. Nhưng sau tất cả thì việc "nấu lẩu thập cẩm" và refactor đập đi xây lại cũng là một "thói quen" rồi. Rồi cũng chọn được cái phù hợp nhất. Quan trọng là tồn tại được tới khi có các engineer cứng vào. Khi anh đã đúng, không ai quan tâm quá khú a đã sai như nào

**4.1. Bổ sung của anh Thai Le**

Database call i/o rất là chậm. Nên cache tầng business server vẫn lầ rất có lợi. Chưa tính network i/o RTT.
