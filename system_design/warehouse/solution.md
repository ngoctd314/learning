<p style="font-size: 38px; font-weight: 700">GHTK challenge: Đếm trên tập lớn</p>

[[TOC]]

## TLDR;

Bài viết này đưa ra 2 solutions tạm gọi là: single instance, multi instances.

Cả 2 solution đều đáp ứng được nhu cầu của bài toán

**Single instance (vertical scale)**

+ MySQL, Golang

+ Với cấu hình server có 12 **physical CPUs** khoảng 0.6s, server có 24 **physical CPUS** khoảng 0.3s

**Multi instances (horizontal scale)**

+ MySQL, Golang, gRPC (hoặc REST)

+ Với cấu hình 5 instances, mỗi instance 6 **physical CPUs** thời gian tính toán khoảng 0.4 s

Vì logic tính toán khá nhiều nên thời gian xử lý phụ thuộc rất nhiều vào CPUs. Cả 2 phương án trên đều có thể giảm thời gian tính toán bằng việc cung cấp thêm CPUs cho mỗi instance. Có thể đạt tới 0.1 - 0.2s nếu bỏ nhiều resource.

## Công cụ sử dụng

- Golang, MySQL, bitmap, bit compression, concurrency, gRPC 
- Redis, bloom-filter

## Mô tả lại bài toán

+ A1 gồm ~5000 số integer (unique) ∈ [1, 1e8]
+ A2 gồm ~5000 số integer (unique) ∈ [1, 1e8]<br>
...
+ An gồm ~5000 số integer (unique) ∈ [1, 1e8]
+ n ∈ [5000, 10000]

**Output:** A1 ∪ A2 ∪ ... An-1 ∪ An

## Yêu cầu chính

+ Đếm: A1 ∪ A2 ∪ ... An-1 ∪ An
+ Thêm: relate_item vào Ai (i ∈ [1, 1e8])
+ Xóa: relate_item j từ Ai (i,j ∈ [1, 1e8])
+ **Count phải nhỏ hơn 1s**

## Yêu cầu khác

+ Triển khai được trên phần cứng phổ thông
+ Scalable, High availability

## Cấu trúc dữ liệu bitmap và thuật toán nén

+ Bài toán đếm phần tử chung giữa 2 **tập hợp các số nguyên**

Có thể dùng nhiều cách để biểu diễn tập hợp như hash, set, array...

Tuy nhiên bài toán này có đặc tính là rất lớn => các cấu trúc trên là không hợp lý.

Để biểu diễn tập hợp số nguyên, ta còn có một cách khác là dùng bitmap.

Nếu max(Ai) = n (phần tử lớn nhất trong tập hợp bằng n). Thì ta có thể sử dụng n bit để biểu diễn tập hợp đó.

Phép count: A1 ∪ A2 ∪ ... An-1 ∪ An

Được biến đổi thành phép đếm số 1 trong phép OR bit: A1 | A2 | ... An-1 | An

Ví dụ:

+ Tập hợp A = {1,2,3,4,5}, B = {1,3,5}

=> có thể dùng 5 bits để biểu diễn: 

A = 11111 <br>
B = 10101 <br>
A | B = 11111<br>
Count1(A|B) = 5 <br>

+ Tập hơp A = {1,2,3,4,10}, B = {1,5,7}

=> có thể dùng 10 bits để biểu diễn:

A = 1000001111<br>
B = 0001010001<br>
A | B = 1001011111<br>
Count1(A|B) = 8

+ Tập hợp A = {1,2,3}

bits(A) = 111

Để thêm số 4 vào tập hợp, sử dụng phép OR bit: 

bits(4) = 1000

A | bits(4) = 1111 = {1,2,3,4}

Bản chất việc thêm 1 phần tử vào 1 tập hợp chính là union 2 tập hợp.

+ Tập hợp A = {1,2,3}

Để bỏ 1 số 1 ra khỏi tập hợp, sử dụng phép AND bit:

bits(1) = 001

A & bits(1) = 110 = {2,3}

=> Với cách biểu diễn theo bit, ta có thể count, thêm phần tử, xóa phần tử...

Tuy nhiên bài toán ban đầu max(Ai) = 100M => biểu diễn theo cách này thì cần tới 100M bits => con số không tưởng.

=> Câu hỏi: Có cách nào để làm cho 100M này nhỏ lại không?

Nhìn lại bài toán thì mỗi item chỉ liên quan đến khoảng 5000 items khác => nếu biểu diễn tập A dưới dạng bitmap với 100M bits thì bitmap này chỉ toàn là số 0 (100M - 5000 là số 0).

=> Có cách nào để nén lại không?

Bài toán ban đầu được đưa thành bài toán nén bit.

Có nhiều thuật toán (thuộc lớp run-length-encoded bitmaps) để nén bit.

+ Oracle's BBC
+ WAH (variation on BBC)
+ EWAH

Các thuật toán này có một điểm yếu là không random access. Nếu muốn kiểm tra một giá trị có ở trong set hay không, ta phải bắt đầu từ vị trí đầu tiên, sau đó giải nén. Điều này có nghĩa là khi intersect hoặc union trên một set rất lớn, ta phải giải nén toàn bộ set => không đáp ứng nhu cầu của bài toán.

Để xử lý phần này, đề xuất sử dụng roaring. Roaring chia dữ liệu thành các khối gồm 2^16 số nguyên. Trong một đoạn, nó có thể sử dụng bitmap không nén, một danh sách số nguyên đơn giản hoặc danh sách các lần chạy. Dù nó sử dụng định dạng nào, tất cả chúng đều cho phép kiểm tra sự hiện diện của bất kỳ một giá trịnào một cách nhanh chóng (ví dụ: bằng tìm kiếm nhị phân). Kết quả cuối cùng là Roaring có thể tính toán nhiều thao tác nhanh hơn nhiều so với các định dạng được mã hóa theo thời lượng chạy như WAH, EWAH...

Xem chi tiết tại: https://roaringbitmap.org/

## Basic approach: Single instance

Phần này đề xuất phương án dễ dàng implement với:

+ 1 service (viết bằng Golang)
+ 1 database (dùng MySQL)

### Thiết kế Database 

relate table

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

**Vì data size khả lớn, cần có đánh partition theo id (PRIMARY KEY) khoảng 1M / 1 partition.**

**+ SELECT**

Lấy ra tất cả bitmap từ danh sách các items (theo đề bài thì là từ 5K -> 10K).

```sql
SELECT bitmap FROM relate WHERE id IN (?);
```

**+ INSERT**

```sql
INSERT INTO relate (bitmap) VALUES (?); 
```

**+ UPDATE**

```sql
UPDATE relate SET bitmap = ? WHERE id = ?;
```

Bảng relate chỉ có 2 trường: id, bitmap. Số lượng ids của 1 kho nằm trong khoảng 5000 -> 10000. Query lấy 10K record theo PRIMARY KEY.

### Cost estimate

**Database**

Size của một record trong bảng relate

|column|type|size|
|-|-|-|
|id|unsigned int|4 bytes|
|bitmap|blob|L + 2 bytes, L < 2^16, đã tính thử max = 30KB|

có 100M records = (4 + 30000)*1e8 bytes ~ 30e8 KB ~ 30e5 MB ~ 30e2 GB ~ 3 TB.

### Scale

**Benchmark 10M dữ liệu trên máy cá nhân (single instance) với chỉ 6 physical CPUs, MySQL và server chạy trên cùng 1 máy**

|Items của kho X|Thời gian thực hiện|< 1s|
|-|-|-|
|1000|135.437861 ms|OK|
|5000|565.390196 ms|OK|
|8000|875.923163 ms|OK|
|10000|1.119023969 s|FAIL (>1s)|

Để tăng tốc độ xử lý cần chạy trên server có nhiều CPUs hơn (vertical scale). Với khoảng 24 physical cores thì dự kiến đếm distinct 10K items có thể đạt 0.2 => 0.3s. Lý do có con số này là phần code đã apply chia nhỏ câu query theo chunk. Mỗi CPUs sẽ xử lý query 100 ids và thực hiện tính toán. Với 1000 ids chia cho 6 CPUs thời gian trung bình chỉ là 130ms. Lưu ý ở đây yêu cầu physical CPUs. Vì các tác vụ đa phần là tính toán, logical CPUs không có nhiều tác dụng trong xử lý tác vụ liên quan đến tính toán trên CPUs.

![alt](./assets/vertical-scale.png)

Phần bottleneck còn có thể xảy ra ở bước query database. Có thể áp dụng đọc trên nhiều db slave. Và xử lý concurrency kết quả trả về từ DB.

![alt](./assets/read-slave.png)

Phần data model cũng chỉ là logic từ 1 id kiểu unsign integer lấy ra mảng bytes (bitmap). Có thể dùng hệ distributed database khác để đảm bảo tính scale.

Về phần Or trên bitmap (implement roaring) thì không có vấn đề gì về hiệu năng.

## Multi instance approach

Phương án single instance sẽ gặp giới hạn về phần cứng vì số lượng CPU trên 1 máy là hữu hạn. Phương án này đưa ra cơ chế cho phép xử lý song song trên nhiều máy. Tuy nhiên sẽ thêm một chút phức tạp và có overhead của việc giao tiếp qua network. Để giám tối đa độ trễ trong việc giao tiếp giữa các service, đề xuất dùng gRPC, tuy nhiên với REST hay cơ chế giao tiếp khác thì implement cũng tương tự, không quá quan trọng.

### Thiết kế database

Như phần trên.

### Scale

#### 1 aggregator, n workers

![alt](./assets/single-aggregator.png)

Aggregator đầu tiên sẽ chia đều lượng ids cho các workers mong muốn sẽ là 2000 ids cho 1 worker. Implement của worker là như phần single instance. Tuy nhiên ở bước cuối cùng, trả về bitmap là Or(ids) chứ không tính trực tiếp ra số lượng phần tử của tập hợp. Bởi vì: 

Ban đầu tập hợp có n ids, chia thành k phần. Kết quả đúng phải là lưc lượng(k1 | k1 ... | kn). Chứ không phải lực lượng k1 + lực lượng k2 ... + lực lượng kn. Do đó worker trả về k1 = bitmap(id1) | bitmap(id2)... | bitmap(idn/k)

Sau đó aggregator sẽ thực hiện phép orBitMap(k1, k2, ..., kn)

#### n aggregators, n workers

Giải pháp trên có bottleneck khi chỉ có một aggregator. Có thể thêm aggregator + load balancer để điều phối requests.

![alt](./assets/multi-aggregator.png)

## Các cách tối ưu cho phần lưu trữ database khác

Nhắc lại tính chất của phép OR trên tập ids n phần tử: id1 | id2 ... | idn = (id1 | id2) | (id3 ... | idk) ... | (idk | idn)

Từ tính chất đó, ta có thể chia database thành các cách như sau:

### Partition

Cách đơn giản nhất là partition chia db ban đầu thành k partition, lưu ý là không nên chia quá nhỏ vì query 10K phần tử trong 100M chia khoảng 10 partition là đủ.

### Chia thêm cột

**1. Thiết kế ban đầu**

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

bitmap sẽ chứa thông tin relate của item với id hiện tại, bitmap có khoảng 5000 phần tử (do 1 item relate tới khoảng 5000 items), và range của bitmap từ 1 -> 100M.

Câu count được tính như sau: 

```txt
count(id1.bitmap | id2.bitmap ... | idn.bitmap)
```

**2. Chia thêm cột như sau**

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap1|blob|represent set of related item|
|bitmap2|blob|represent set of related item|
|bitmap3|blob|represent set of related item|
|bitmap4|blob|represent set of related item|
|bitmap5|blob|represent set of related item|

Khi này bitmap1 sẽ chứa thông tin relate item với các item khác nằm trong khoảng 1 -> 20M<br>
Khi này bitmap2 sẽ chứa thông tin relate item với các item khác nằm trong khoảng 20 -> 40M<br>
...
Khi này bitmap5 sẽ chứa thông tin relate item với các item khác nằm trong khoảng 80M -> hết

Câu count được tính như sau: 

```txt
count(id1.bitmap1 | id2.bitmap1 ... | idn.bitmap1) +
count(id1.bitmap2 | id2.bitmap2 ... | idn.bitmap2) +
... +
count(id1.bitmapn | id2.bitmapn ... | idn.bitmapn)
```

### Chia thêm bảng

**1. Bảng ban đầu**

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

Với id thuộc [1,100M]

**2. Ta có thể chia thành 5 bảng như sau**

+ relate20M table

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

Với id thuộc [1,20M]

+ relate40M table

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

Với id thuộc [20M,40M]

...

+ relate100M table

|Column|Data type|Description|
|-|-|-|
|id|int unsigned not null|item id + PRIMARY KEY|
|bitmap|blob|represent set of related item|

Với id thuộc [80M, x]

Câu count được tính như sau: 

```txt
count(
    (tbl1.id1.bitmap | tbl1.id2.bitmap ... | tbl1.idn.bitmap) |
    (tbl2.id1.bitmap | tbl2.id2.bitmap ... | tbl2.idn.bitmap) |
    ... |
    (tbln.id1.bitmap | tbln.id2.bitmap ... | tbln.idn.bitmap)
)
```

### Tách database

Giống như cách chia bảng nhưng mỗi bảng bây giờ sẽ nằm ở một db riêng. Tuy nhiên cách này sẽ phải xử lý distributed transaction trong trường hợp update.

## Xử lý hot key và bloom-filter

Nếu xuất hiện những item là hot key trong hệ thống thì có thể cache lại, giảm tải cho db. Tuy nhiên vì không gian rất lớn (100M) items, mỗi lần query lại lấy ra thông tin của 10000 items. Nếu lần nào cũng check cache thì lại là bottleneck. Yêu cầu là có một cách nào đó tốn chi chí rất thấp mà để có thể đưa ra quyết định có vào cache để lấy hay không. 

Bloom-filter là một cấu trúc dữ liệu như thế. Bloom-filter trả lời câu hỏi sự tồn tại của một phần tử trong tập hợp.

Phần tử A có trong tập X hay không. Kết quả của bloom-filter sẽ là:

+ Chắc chắn không có
+ Có thể có

Trong case chắc chắn không thì đưa cho worker xử lý như bình thường

Còn case có thể có thì kiểm tra cache.

## Kiểm thử

**1. Để kiểm thử tính chính xác của thuật toán, cần sinh ra một file data.txt với định dạng như đề bài**

1 2 3 4<br>
2 1 3 5<br>
3 1 2<br>
4 2 5<br>
5 1 4<br>

Với cột đầu tiên là id, các cột tiếp theo là các item liên quan

**2. Tạo database cho việc test, thông số cấu hình db điền trong file .env**

**3. Tạo table relate:**

+ Thủ công

```sql
CREATE TABLE relate (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    bitmap blob
);
```

**4. Sau đó chạy lệnh:**

```bash
./cmd/insert_from_file
```

Lệnh này sẽ insert data từ file data.txt vào database.

**5. Đếm distinct relate_item bằng lệnh:**

Đếm unique của relate items có id 1,2,3,4

```bash
./cmd/count 1,2,3,4
```
