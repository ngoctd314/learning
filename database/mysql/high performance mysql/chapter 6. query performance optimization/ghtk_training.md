# GHTK training

Parser
Resolver (Prepare)

**Optimizer**

- Logical transformation: transform alge
- Cost-based optimizer
- Plan refinement

Query execution
Storage engine

```sql
SELECT * FROM t1, t2 WHERE
    t1.a = t2.a AND t2.a = 9 AND (NOT (t1.a > 10 OR t2.b > 3) OR (t1.b = t2.b + 7 AND t2.b = 5));

-- Negation elimination
SELECT * FROM t1, t2 WHERE
    t1.a = t2.a AND t2.a = 9 AND (t1.a <= 10 AND t2.b <= 3 OR (t1.b = t2.b + 7 AND t2.b = 5));

-- Equality/const propagation
SELECT * FROM t1, t2 WHERE
    t1.a = 9 AND t2.a = 9 AND (t1.a <= 10 AND t2.b <= 3 OR (t1.b <= 5 + 7 AND t2.b = 5));

-- Evaluate const expressions
SELECT * FROM t1, t2 WHERE
    t1.a = 9 AND t2.a = 9 AND (t2.b <= 3 OR (t1.b = 12 AND t2.b = 5));
```

## Cost-based Optimizer

### Summary

- SQL is expect of dev, not execution plan of query statement
- Query statement must run over multiple statistic, transform and optimize then execute. Statement execution order maybe different with order describe in SQL
- Execution plan from optimizer not always fastest.

##

### Chuẩn đoán lâm sàng: Query Cost

- Query run time: Đơn giản, khó bắt đúng bệnh và chữa tận gốc vấn đề

- Row Sent, Row Examined: Tối ưu được phần lớn các câu truy vấn. Chưa hiểu đúng bản chất kế hoạch thực thi

- Cost-based Optimization: Hiểu đúng bản chất Database thực thi truy vấn. Khá khó hiểu với người mới

### Query Cost

**Cost-based:**

- Index access
- Table scan
- Index scan, range, scan, ref access
- Join order
- Subquery

**Cost model**

Cost formulas: Access method, Join, Subquery

Cost constants: CPU, IO

| Cost                                | MySQL 5.7 | MySQL 8.0 |
| ----------------------------------- | --------- | --------- |
| Read a random disk page             | 1.0       | 1.0       |
| Read a data page from memory buffer | 1.0       | 0.25      |
| Evaluate query condition            | 0.2       | 0.1       |
| Compare keys/records                | 0.1       | 0.05      |

**Metadata**

- Record, index size
- Index information
- Uniqueness

**Statistics**

- Table size
- Cardinality
- Range estimates
- Histograms

```sql
CREATE TABLE `package_addresses` (
    `id` BIGINT(20) NOT NULL AUTO INCREMENT,
    `package_order` BIGINT(20) NOT NULL,
    `type` ENUM('pickup','deliver','return') NOT NULL,
    `tel` VARCHAR(50) DEFAULT NULL,
    `first_address` VARCHAR(500) DEFAULT NULL,
    `cart_id` CHAR(36) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UNIQ` (`package_order`, `type`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

```sql
CREATE TABLE `carts` (
    `id` CHAR(36) NOT NULL,
    `type` TINYINT(4) NOT NULL COMMENT '1:picking, 2:delivering, 3:returning',
    `alias` VARCHAR(100) NOT NULL,
    `order` BIGINT UNSIGNED DEFAULT NULL,
    `first_cod_id` CHAR(36) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `IDX UNIQ` (`order`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8mb4_unicode_ci;
```

```sql
SELECT pa.`package_order`, pa.`first_address`, c.`alias`, c.`first_cod_id`
FROM `package_address` pa
INNER JOIN carts C ON c.id = pa.cart_id
WHERE pa.package_order = 1510775540
```

**select_type:** loại select dùng trong câu truy vấnot

- SIMPLE: là một câu SELECT cơ bản, không có subqueries, UNION
- PRIMARY: là câu SELECT ngoài cùng của một phép JOIN (Outermost SELECT)
- DERIVED: Query nằm bên trong FROM
- SUBQUERY: Query đầu tiên trong subquery, không phụ thuộc vào query khác. Query này được execute 1 lần và cache lại
- DEPENDENT SUBQUERY: Subquery phụ thuộc vào query nằm ngoài nó
- UNCACHEABLE SUBQUERY: Query không thể cache lại được
- UNION: Query là câu SELECT thứ hai của lệnh UNION
- DEPENDENT UNION: Trong subquery có union, và subquery đó thuộc loại DEPENDENT Subquery
- UNCACHEABLE UNION

**type:** Join type, table scan type

- const: so sánh bằng trên cột Primary/Unique Key với 1 hằng số
- system: bảng có duy nhất 1 row
- eq_ref: Join/so sánh 1-1 trên cột Primary Key/Unique key với 1 cột ở bảng trước đó. (Duy nhất 1 row)
- ref: Join/So sánh dự trên key (Không phải Unique key, phép join 1-n)
- ref_or_null: Giống với ref và kiểm tra thêm điều kiệu null
- fulltext: join sử dụng full text index
- index_merge: sử dụng index merge
- unique_subquery: thay thế eq_ref với IN subquery
- index_subquery: tương tự unique_subquery với trường hợp subquery sử dụng non-unique indexes.
- range: sử dụng index để range scan bảng
- ALL: full table scan
- index: Full index scan

**possible_keys:** Index có thể sử dụng được

**key_len:** Độ dài key (index leftmode prefix) được sử dụng (bytes)

**rows:** số rows examined phải scan

**filtered:** estimate % số row còn lại khi filter bởi điều kiện WHERE

**extra**

## Explain format json

- Lợi ích:

Cho chúng ta biết chi tiết kế hoạch thực thi của câu truy vấn.

Các chỉ số giúp chúng ta hiểu về quá trình thực thi truy vấn và giúp tìm cách tối ưu câu truy vấn.

- Hạn chế:

Không nói cho chúng ta biết bản chất vấn đề. Tại sao kế hoạch thực thi đó lại tối ưu và được lựa chọn

Không có thông tin về triggers, store functions

Các con số thống kê chỉ là ước lượng và có thể outdate

Không nói lên tất cả vấn đề

## Optimizer trace

Explain: Trả lời câu hỏi HOW, diễn giải kế hoạch thực thi được chọn

- Kế hoạch thực thi câu truy vấn như thế nào

Optimizer trace: Trả lời câu hỏi WHY

- Tại sao kế hoạch đó được chọn, có những kế hoạch nào khác?
- Quá trình thực thi chi tiết và nhiều thông tin khác.

### Where condition

- Đơn giản hóa logic: bộ optimizer không phải bao giờ cũng hiệu quả. Hãy biến đổi logic về đơn giản nhất, chuyển các tính toán logic sang vế phải
- Đánh index hiệu quả

- Leftmode prefix
- Index (const, range)
- Multi-column index
- Index condition pushdown
- Index merge

### Join

#### Inner join vs left join

- Chỉ định thứ tự Join bằng left join
- Chú ý nếu left join có mệnh đề where trên bảng bên phải thì left join ~ inner join

SELECT \* FROM t1 LEFT JOIN t2 ON (column1) WHERE t2.column2=5;

#### NESTED loop join

- Thuật toán join được implement ở hầu hết database
- Đặt bảng có điều kiện filter tốt sang trái

#### Sử dụng Count(\*) thay COUNT(col)

Sử dụng COUNT(_) nếu muốn count số records trả về. Hãy dùng COUNT(_) để tránh nhập nhằng NULL value

#### Tính xấp xỉ

Sử dụng EXPLAIN để tính xấp xỉ số row trả về

SHOW TABLE STATUS estimate số row của bảng

#### Kỹ thuật Caching

### ORDER BY, GROUP BY

- Sử dụng index cho các toán tử ORDER BY, GROUP BY

* Đánh index với cột trong mệnh đề ORDER BY, GROUP BY để hạn chế SORTING
* Với MySQL 8.0 đánh Descending indexes cho các toán tử ORDER BY ASC, DESC custom

- ORDER BY NULL

* Trong các toán tử GROUP BY nếu không cần sắp xếp kết quả, hãy thêm ORDER BY NULL để hạn chế SORTING

- Sử dụng các kỹ thuật khác

Bản thân các database OLTP không thiết kế phù hợp cho các câu truy vấn tổng hợp, thống kê.

Hãy sử dụng database OLAP, hoặc tính toán trước kết quả qua các kỹ thuật CDC, tmp table, view...

### LIMIT, OFFSET

**LIMIT dữ liệu trả về**

- Nếu không cần phải lấy toàn bộ dữ liệu, hãy LIMIT để hạn chế số lượng records trả về
- Không SELECT \*, chỉ select các cột cần thiết

**Cẩn thận với LIMIT OFFSET**

- Câu truy vấn LIMIT 10 OFFSET 10000 tương đương với việc database phải truy cập qua 10010 records để lấy 10 records cuối cùng
- Với các màn hình phân trang, các màn cuối thường sẽ rất chậm nếu sử dụng LIMIT OFFSET

=> Chuyển qua phân trang WHERE trên KEY

```sql
SELECT id, username, fullname
FROM users LIMIT 10 OFFSET 1000;

SELECT id, username, fullname
FROM users WHERE id >= 1000 AND id < 1010;

SELECT id, usernamee, fullname
FROM users
ORDER BY id LIMIT 10 -- last id = 123

SELECT id, username, fullname
FROM users WHERE id > 123
ORDER BY id LIMIT 10;
```

**Tách truy vấn**

## Union

- MySQL mặc định sẽ tạo bảng tạm (mem hoặc disk) để gộp kết quả của các câu truy vấn. Chi phí cho UNION khá lớn.
- Hãy LIMIT kết quả trả về nhất có thể trước khi UNION

```sql
( SELECT username, fullname FROM user ORDER BY username LIMIT 20 )
UNION ALL
( SELECT username, fullname FROM ightk_users ORDER BY username LIMIT 20 )
LIMIT 20;
```

- Nếu không cần distinct kết quả trả về, hãy sử dụng UNION ALL để tránh MySQL phải thêm thao tác so sánh, loại bỏ giá trị trùng lặp.

## Distinct, Subquery

**Sử dụng EXISTS để tránh DISTINCT**

Trong một số trường hợp cần DISTINCT kết quả, có thể chuyển qua dùng EXIST để tránh việc phải tạo bảng tạm.

**Convert Subquery**

Với MySQL 5.7 trở về trước, convert subquery WHERE IN (SELECT ... FROM ...) sang EXIST hoặc chuyển về JOIN.

Hạn chế dependency subquery.
