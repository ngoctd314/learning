# Database

## 

**1. Redis server sử dụng giao thức nào sau đây?**

A client connects to a Redis server by creating a TCP connection to its port (the default is 6379).

**2. [MySQL] Nếu cột Extra khi thực thi câu lệnh explain statement hiển thị: "Using temporary; Using filesort	có	ý	nghĩa	là?	(Có	thể	có	nhiều	đáp	án)**

Trong câu truy vấn có thể có các toán tử GROUP BY, ORDER BY, DISTINCT.

**3. Phantom Read trong Database đề cập tới hiện tượng gì?**

Khi 2 truy vấn giống y hệt nhau được thực thi nhưng kết quả trả về quả query thứ 2 lại khác với query đều tiên. (Nhìn thấy một bản ghi mới (phantom) hoặc mất đi vài bản ghi(phantom))

**4. Trường hợp có nguy cơ cao về xung đột cập nhật vào một bản ghi trong database, ta nên dùng cách nào trong các cách sau?**

Set Isolation level, select for update, lock row

**5. Chữ "d	trong	hình	bên	cho	biết	điều	gì?	\nhttps://cache.giaohangtietkiem.vn/d/907d4e5246d621d4fd8b5f67bddfae50.png**

**6. Cài đặt nào sau đây trong MySQL có thể dẫn tới kết quả sai? \nhttps://cache.giaohangtietkiem.vn/d/9cd0225eb3d274d24c0e67bb520fa43f.png?width=435&height=450** 

**7. Trong InnoDB, mặc định có bao nhiêu pages được gom lại trong một block (còn gọi là extent)?**

An InnoDB data file is a sequence of equal sized pages. These pages are grouped into extents and segments. One extent is a contiguous sequence of pages in an idb file. The number of pages belonging to an extent depends on the page size.

Page size default is 16 KB so Extent Size in Pages is 64 and Extent size in MB = 1 MB

An extent is the basic unit of file space allocation in InnoDB.

**8. Câu lệnh nào liên quan đến việc kiểm soát giao dịch (transaction) trong SQL?**

BEGIN TRAN T1;
COMMIT TRAN T1;
ROLLBACK TRAN T1;

**9. Cho bảng users có cấu trúc:\nCREATE TABLE `users` (\n  `id` int(11) NOT NULL,\n  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Tên người dùng',\n  `group_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'nhóm tài khoản, 1 user có thể thuộc nhiều nhóm',\n  `birthday` date DEFAULT NULL COMMENT 'Ngày sinh',\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;\n\nSample data:\nid | name | group_id | birthday\n1 | Phan Thiên Hạo | 1,2 | 2001-04-06\n2 | Trần Thiên Di | 1 | 1999-10-02\n\nBảng trên đang được thiết kế ở dạng chuẩn nào?**

id name group_id birthday 
1  H    1,2      1999-01-01

=> NON NF
 
**10. CREATE TABLE t1 (i1 TINYINT, i2 TINYINT UNSIGNED); \nINSERT INTO t1 (i1, i2) VALUES(256, 256); \nSelect * from t1;**

**11. Isolation level mặc định trong MySQL InnoDB là gì?**

REPEATABLE READ

**12. MySQL: Phát biểu nào sau đây đúng về charset utf8 và utf8mb4?**

UTF-8 is a variable-length encoding. In the case of UTF-8, this means that storing one code point requires one to four bytes. However, MySQL's encoding called "utf8" (alias of "utf8mb3") only stores a maximum of three bytes per code point.

utf8mb4 character set uses a maximum of four bytes per character supports supplemental characters

**13. Nếu cột Extra khi thực thi câu lệnh explain statement có giá trị là: "Using temporary; Using filesort	nghĩa	là?**

GROUP BY, DISTINCT, ORDER BY

**14. Dirty Read trong Database là hiện tượng gì?**

Read data from another transaction

**15. Dữ liệu của Table được lưu tại đâu trong những thành phần dưới đây của MySQL-InnoDB?**

InnoDB stores its tables and indexes in a tablespace, which may consist of several files.

**16. Đâu không phải là 1 format của binlog?**

Binary log contains "events" that describe database changes such as table creation operations or changes to table data. It also contains events for statements that potentially could have made changes. The binary log also contains information about how each statement took that updated data. The binary log has two important purposes: 

- For replication, the binary log on a replication source server provides a record of the data changes to be sent to replicas. The source sends the information contained in its binary log to its replicas, which reproduce those transactions make the same data changes that were made on the source.

- Certain data recovery operations require use of the binary log. After a backup has been restored, the events in the binary log that were recorded after the backup was made are re-executed.

The binary log is not used for statement such as SELECT or SHOW that do not modify data. To log all statements, use the general query log.

Running a server with binary logging enabled makes performance slightly slower. However, the benefits of the binary log in enabling you to set up replication and for restore operations...

The server uses several logging formats to record information in the binary log:

- Replication capabilities in MySQL originally were based on propagation of SQL statements from source to replica. This is called statement-based logging. You can cause this format to be used by starting the server with --binlog-format=STATEMENT
- In row-based logging (the default), the source writes events to the binary log that indicate how individual table rows are affected. --binlog-format=ROW
- A third option is also avaiable: mixed loggin. With mixed logging, statement-based logging is used by default, but the logging mode switches automatically to row-based in certain cases. --binlog-format=MIXED.

**17. Đâu là phương án lựa chọn NoSQL database phù hợp nhất cho các bài toán sau:\n1. Schema đơn giản, đọc/ghi tốc độ cao, it có update, có thể scale, truy vẫn không phức tạp, không liên quan đến nhiều khoá\n2. Schema linh hoạt, truy vấn phức tạp, định dạng json/bson hoặc xml, hiệu năng cao, cân bằng giữa đọc/ghi\n3. Kích thước dữ liệu lớn, trích xuất dữ liệu theo cột, không có mẫu truy vấn đặc biệt, mức độ tổng hợp (Aggregations) cao\n4. Các ứng dụng yêu cầu truyền tải giữa các điểm dữ liệu, khả năng lưu trữ thuộc tính cũng như mối quan hệ của từng điểm dữ liệu (Data point)**

Schema đơn giản, đọc/ghi tốc độ cao, ít update, truy vấn không phức tạp, không liên quan đến nhiều khóa.

**18. Giả sử có 2 bảng có cùng cấu trúc (id, age, name) có 10 triệu bản ghi giống hệt nhau. Bảng T1 có index IDX(age) bảng T2 thì không với 2 queries:\nQ1: Select * from T1 where age * 10 > 50\nQ2: Select * from T2 where age * 10 > 50\nGiả sử t1 và t2 là thời gian thực thi 2 câu query Q1 và Q2**

t1 as same as t2 because where cond is not atomic

**19. Phát biểu đúng về TRUNCATE trong SQL?**

The TRUNCATE command is used to remove all of the rows from a table, regardless of whether or not any conditions are met and resets the table definition. It is a DDL (Data Definition Language) command. The table space is completely freed from the memory.

**20. SQL-DDL-DQL-DML-DCL-TCL**

https://www.geeksforgeeks.org/sql-ddl-dql-dml-dcl-tcl-commands/

**21. Trong MySQL điều nào sau đây là đúng khi nối về khóa chính (Primary key)?**

**22. ON DELETE CASCADE mang ý nghĩa gì?**

**23. Port mặc định của MySQL là bao nhiêu?**
3306

**24. RDBMS là viết tắt của từ gì trong tiếng Anh?**

Relational Database Management Systems

**25. Số bytes dùng để lưu trữ kiểu dữ liệu SMALINT trong MySQL?**

TINYINT 1 bytes, SMALINT 2 bytes, MEDIUMINT 3 bytes, INT 4 bytes, BIGINT 8 bytes

**26. Trong Mysql toán tử like ký tự % thể hiện điều gì?**

**27. Trong SQL, làm thế nào để chọn tất cả các cột dữ liệu trong bảng Persons:**

SELECT column_name FROM information_schema.columns
WHERE table_schema = 'Persons'
ORDER table_name, ordinal_position

SELECT `COLUMN_NAME` FROM `COLUMNS` WHERE `TABLE_NAME` LIKE '%account_permission%' AND `TABLE_SCHEMA` = 'auth';

**28. Cho đồ thị G đơn vô hướng với 100 đỉnh. Số lượng cạnh tối đa trong đồ thị G mà G là không liên thông?**

<p>Thời gian thực hiện của thuật toán Bellman Ford?</p>										
<p>Thuật toán Dijkstra's không sử dụng được trong trường hợp nào?</p>										
<p>Chuyển đổi biểu thức trung tố sang biểu thức hậu tố?\n A - B / (C * D)</p>										
<p>Heap là một cấu trúc dữ liệu được implement từ?</p>										
<p>Thuật toán nào sau đây được sử dụng cho Fibonacci?</p>										
<p>Thuật toán sắp xếp nào sau đây là ổn định?</p>										
<p>Thuật toán search nào sau đẩy có thể sử dụng đệ quy?</p>										
<p>Tính giá trị của hàm F(3)?\n F(0) = 3; F(n) = 2F(n-1) + (n-1)^2</p>										
<p>Binary search hiệu quả trên?</p>										
<p>Heap luôn là?</p>										
<p>Min Heap là 1 cấu trúc cây, phần tử bé nhất xuất hiện ở?</p>										
<p>Một cấu trúc cây có N đỉnh, thì chứa bao nhiêu cạnh?</p>										
<p>Số lần swap cần thiết để sắp xếp dãy sau bằng thuật toán bubble sort?\n 8, 22, 7, 9, 31, 19, 5, 13</p>										
<p>Thuật toán quicksort sử dụng tư tưởng thiết kế nào sau đây?</p>										
<p>Hệ thống gửi mail của GHTK sử dụng Kafka làm queue. Kafka topic có 2 patition, với 3 consumer xử lý. Vào thời điểm cao tải, hệ thống thường xuyên bị hiện tượng gửi mail chậm do lượng job tồn trong queue lớn. Đâu là phương án tối ưu nhất để giảm số lượng mail tồn trong queue?</p>										
<p>Trong nguyên lý CAP của distribute system, Redis thuộc nhóm nào sau đây?</p>										
<p>Hệ thống của bạn chỉ có 1 API có kết nối đến MySQL để xử lý dữ liệu, sau khi deploy một tính năng mới, phía application xảy ra lỗi “Too many connections”. Bạn sẽ có thể thực hiện những biện pháp nào để fix vấn đề?</p>										
<p>Khi server (tầng application) của 1 services bị hiện tượng OOM (Out of memory) nguyên nhân có thể do đâu? (Chọn nhiều đáp án)</p>										
<p>Đâu là những thước đo hữu ích cho Performance Testing?</p>										
<p>Phát biểu nào sau đây là đúng khi nói về việc thêm header "Connection: keep-alive	trong	HTTP?	(Chọn	nhiều	đáp	án)</p>"				
<p>Câu nào sau đây mô tả sai về cookie và session?</p>										
<p>Trong domain http://www.google.com thành phần nào được gọi là protocol?</p>										
<p>CCU là chỉ số gì khi đo lường performance của 1 hệ thống?</p>										
<p>Chỉ số nào dưới đây được đo bằng MTBF (Mean Time Netween Failures)?</p>										
<p>Đâu không phải là một tiêu chí đo lường performance của 1 services?</p>										
<p>Hệ thống có thể mở rộng để đáp ứng được việc tăng tải được gọi là?</p>										
<p>Phương án rate limit để tránh thất thoát dữ liệu trong lập trình?</p>										
<p>Dữ liệu ở token JWT khi giải mã base64 có thể thấy được các scope, ta có thể sửa scope rồi mã hóa lại token để lấy thêm quyền không? Vì sao?</p>										
<p>Dự án cần lựa chọn phương án để mã hóa giá trị căn cước công dân trong database, phương án mã hóa nào sau đây là hợp lý và an toàn?</p>										
<p>Các phương án nào sau đây có thể làm thay đổi IP public của người gọi request (để giả mạo IP)?</p>										
<p>Lỗ hổng đọc file bất kỳ ở ứng dụng chạy trên linux có thể đọc được các giá trị gì sau đây (ứng dụng không chạy quyền root)?</p>										
<p>Khi xử lý parse file xml cần lưu ý điều gì để tránh lỗ hổng XXE (XML external entity)</p>										
<p>Các biện pháp để tăng cường tính an toàn cho xác thực đa nhân tố là?</p>										
<p>Cho api như sau: GET /my-user-info trả về thông tin của user hiện tại, nhưng khi truyền thêm param ?user_id=123 thì lại thấy trả về thông tin của user 123, đây là lỗ hổng gì?</p>										
<p>Chỉ thị nào sau đây cho phép accept cookie từ trình duyệt trong cross site request</p>										
<p>Các HTTP Header nào sau đây không được tùy biến giá trị bằng javascript trên trình duyệt cài đặt mặc định?</p>										
<p>Thuật toán mã hóa nào dưới đây khác với các thuật toán còn lại?</p>										
<p>Trong trường hợp cần lưu thông tin token của tổ chức tín dụng để thực hiện thanh toán định kì thì thuật toán mã hóa nào là phù hợp?</p>										
<p>Ứng dụng có sử dụng mã hóa với dữ liệu trong database, trong quá trình xử lý và trao đổi dữ liệu thì cần sử dụng key để mã hóa/giải mã, phương án lưu trữ key nào trong các phương án trên là hợp lý nhất?</p>										
<p>Phương án nào để chống hoặc giảm thiểu tấn công vét cạn?</p>										
<p>Phương pháp nào sau đây dùng để chống lỗ hổng XSS (Cross Site Scripting) hợp lý nhất?</p>										
<p>Redis sử dụng cú pháp nào để xóa toàn bộ keys trong 1 DB?</p>										
<p>Sử dụng câu lệnh nào sau đây để cập nhật cấu hình của redis?</p>										
<p>Redis hỗ trợ bao nhiêu loại dữ liệu?</p>										
<p>Giá trị chuỗi có thể lưu trong redis có kích thước tối đa bao nhiêu?</p>										
<p>Số lượng replication factor tối đa cho 1 topic trong 1 cụm có 9 brokers là?</p>										
<p>Câu lệnh để thực hiện tạo 1 topic trên kafka là?</p>										
<p>Đặc trưng của Kafka là?</p>										
<p>Cấu hình nào chỉ định thời gian giữ logs trong segment trước khi bị xóa đi trong kafka?</p>										
<p>Communication port mặc định của Elasticsearch là?</p>										
<p>Inverted Index được xây dựng và cập nhật bằng cách sử dụng phân tách các word trong documents? Quá trình phân tách các từ gọi là gì?</p>										
<p>Ý nghĩa fuzzy query trong Elasticsearch là gì?</p>										
<p>NRT search trong Elasticsearch là gì?</p>										
<p>Trong ELK, ý nghĩa chữ L là gì?</p>										
<p>Kỹ thuật checksum trong Git sử dụng là gì?</p>										
<p>Công cụ nào sau đây để thực hiện automation testing trên ứng dụng web?</p>										
<p>Cú pháp git checkout dùng để làm gì?</p>										
<p>Continuous Integration gồm những bước nào?</p>										
<p>Mở rộng hệ thống bằng cách bổ sung thêm nhiều nodes được gọi là?</p>										
<p>Trong nguyên lý CAP của distribute system, chữ C thể hiện cho?</p>										
<p>HTTP status nào sau đây là sai?</p>										
<p>HTTP method nào được sử dụng để tạo 1 mới resource?</p>										
<p>Cấu hình tham số URL RESTful để tìm kiếm 1 đơn hàng thông qua order như thế nào?</p>										
<p>Password của users được lưu trữ ở đâu?</p>										
<p>Mô hình mạng máy tính nào thường được sử dụng trong tòa nhà?</p>										
<p>Trong OSI tầng nào cung cấp dịch vụ cho người dùng?</p>										
<p>Kích thước của địa chỉ MAC là?</p>										
<p>IP nào sau đây thuộc lớp B?</p>										
<p>Facade thuộc nhóm design patterns nào?</p>										
<p>Tính chất nào trong OOP nhằm giảm thiểu việc các class lồng nhau?</p>										
<p>Thành phần private của lớp là?</p>										
<p>Trong lập trình hướng đối tượng các hàm có thể trùng tên nhau gọi là gì?</p>										
<p>Sơ đồ nào sau được biểu diễn hoạt động của hệ thống?</p>										
<p>"Mọi thuộc tính không khóa đều không phụ thuộc bắc cầu vào khóa chính	là	thuộc	tính	của	chuẩn	nào	sau	đây?</p>"		
<p>Các câu lệnh CREATE, DROP, ALTER thuộc thành phần nào của SQL?</p>										
<p>Các câu lệnh SELECT, UPDATE, DELETE thuộc thành phần nào của SQL?</p>										
<p>Câu lệnh Alter table có tác dụng gì?</p>										
<p><pre>What will be the output of the following Java code?<br/>class exception_handling <br/>{<br/>public static void main(String args[])<br/>{<br/>try <br/>{<br/>throw new NullPointerException ("Hello");<br/>System.out.print("A");<br/>}<br/>catch(ArithmeticException e) <br/>{<br/>System.out.print("B"); <br/>}<br/>}<br/>}</pre></p>										
<p>Redis server sử dụng giao thức nào sau đây?</p>										
<p>[MySQL] Nếu cột Extra khi thực thi câu lệnh explain <statement> hiển thị: "Using temporary; Using filesort	có	ý	nghĩa	là?	(Có	thể	có	nhiều	đáp	án)</p>"
<p>Phantom Read trong Database đề cập tới hiện tượng gì?</p>										
<p>Trường hợp có nguy cơ cao về xung đột cập nhật vào một bản ghi trong database, ta nên dùng cách nào trong các cách sau?</p>										
<p><pre>Chữ "d	trong	hình	bên	cho	biết	điều	gì?	<br/>https://cache.giaohangtietkiem.vn/d/907d4e5246d621d4fd8b5f67bddfae50.png</pre></p>"		
<p><pre>Cài đặt nào sau đây trong MySQL có thể dẫn tới kết quả sai? <br/>https://cache.giaohangtietkiem.vn/d/9cd0225eb3d274d24c0e67bb520fa43f.png?width=435&height=450</pre></p>										
<p>Trong InnoDB, mặc định có bao nhiêu pages được gom lại trong một block (còn gọi là extent)?</p>										
<p>Câu lệnh nào liên quan đến việc kiểm soát giao dịch (transaction) trong SQL?</p>										
<p><pre>Cho bảng users có cấu trúc:<br/>CREATE TABLE `users` (<br/>`id` int(11) NOT NULL,<br/>`name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Tên người dùng',<br/>`group_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'nhóm tài khoản, 1 user có thể thuộc nhiều nhóm',<br/>`birthday` date DEFAULT NULL COMMENT 'Ngày sinh',<br/>PRIMARY KEY (`id`)<br/>) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;<br/><br/>Sample data:<br/>id | name | group_id | birthday<br/>1 | Phan Thiên Hạo | 1,2 | 2001-04-06<br/>2 | Trần Thiên Di | 1 | 1999-10-02<br/><br/>Bảng trên đang được thiết kế ở dạng chuẩn nào?</pre></p>										
<p><pre>CREATE TABLE t1 (i1 TINYINT, i2 TINYINT UNSIGNED); <br/>INSERT INTO t1 (i1, i2) VALUES(256, 256); <br/>Select * from t1;</pre></p>										
<p>Isolation level mặc định trong MySQL InnoDB là gì?</p>										
<p>MySQL: Phát biểu nào sau đây đúng về charset utf8 và utf8mb4?</p>										
<p>Nếu cột Extra khi thực thi câu lệnh explain <statement> có giá trị là: "Using temporary; Using filesort	nghĩa	là?</p>"								
<p>Dirty Read trong Database là hiện tượng gì?</p>										
<p>Dữ liệu của Table được lưu tại đâu trong những thành phần dưới đây của MySQL-InnoDB?</p>										
<p>Đâu không phải là 1 format của binlog?</p>										
<p><pre>Đâu là phương án lựa chọn NoSQL database phù hợp nhất cho các bài toán sau:<br/>1. Schema đơn giản, đọc/ghi tốc độ cao, it có update, có thể scale, truy vẫn không phức tạp, không liên quan đến nhiều khoá<br/>2. Schema linh hoạt, truy vấn phức tạp, định dạng json/bson hoặc xml, hiệu năng cao, cân bằng giữa đọc/ghi<br/>3. Kích thước dữ liệu lớn, trích xuất dữ liệu theo cột, không có mẫu truy vấn đặc biệt, mức độ tổng hợp (Aggregations) cao<br/>4. Các ứng dụng yêu cầu truyền tải giữa các điểm dữ liệu, khả năng lưu trữ thuộc tính cũng như mối quan hệ của từng điểm dữ liệu (Data point)</pre></p>										
<p><pre>Giả sử có 2 bảng có cùng cấu trúc (id, age, name) có 10 triệu bản ghi giống hệt nhau. Bảng T1 có index IDX(age) bảng T2 thì không với 2 queries:<br/>Q1: Select * from T1 where age * 10 > 50<br/>Q2: Select * from T2 where age * 10 > 50<br/>Giả sử t1 và t2 là thời gian thực thi 2 câu query Q1 và Q2</pre></p>										
<p>Phát biểu đúng về TRUNCATE trong SQL?</p>										
<p>Trong MySQL điều nào sau đây là đúng khi nối về khóa chính (Primary key)?</p>										
<p>ON DELETE CASCADE mang ý nghĩa gì?</p>										
<p>Port mặc định của MySQL là bao nhiêu?</p>										
<p>RDBMS là viết tắt của từ gì trong tiếng Anh?</p>										
<p>Số bytes dùng để lưu trữ kiểu dữ liệu SMALINT trong MySQL?</p>										
<p>Trong Mysql toán tử like ký tự % thể hiện điều gì?</p>										
<p>Trong SQL, làm thế nào để chọn tất cả các cột dữ liệu trong bảng Persons:</p>										
<p>Cho đồ thị G đơn vô hướng với 100 đỉnh. Số lượng cạnh tối đa trong đồ thị G mà G là không liên thông?</p>										
<p>Thời gian thực hiện của thuật toán Bellman Ford?</p>										
<p>Thuật toán Dijkstra's không sử dụng được trong trường hợp nào?</p>										
<p><pre>Chuyển đổi biểu thức trung tố sang biểu thức hậu tố?<br/>A - B / (C * D)</pre></p>										
<p>Heap là một cấu trúc dữ liệu được implement từ?</p>										
<p>Thuật toán nào sau đây được sử dụng cho Fibonacci?</p>										
<p>Thuật toán sắp xếp nào sau đây là ổn định?</p>										
<p>Thuật toán search nào sau đẩy có thể sử dụng đệ quy?</p>										
<p><pre>Tính giá trị của hàm F(3)?<br/>F(0) = 3; F(n) = 2F(n-1) + (n-1)^2</pre></p>										
<p>Binary search hiệu quả trên?</p>										
<p>Heap luôn là?</p>										
<p>Min Heap là 1 cấu trúc cây, phần tử bé nhất xuất hiện ở?</p>										
<p>Một cấu trúc cây có N đỉnh, thì chứa bao nhiêu cạnh?</p>										
<p><pre>Số lần swap cần thiết để sắp xếp dãy sau bằng thuật toán bubble sort?<br/>8, 22, 7, 9, 31, 19, 5, 13</pre></p>										
<p>Thuật toán quicksort sử dụng tư tưởng thiết kế nào sau đây?</p>										
<p>Hệ thống gửi mail của GHTK sử dụng Kafka làm queue. Kafka topic có 2 patition, với 3 consumer xử lý. Vào thời điểm cao tải, hệ thống thường xuyên bị hiện tượng gửi mail chậm do lượng job tồn trong queue lớn. Đâu là phương án tối ưu nhất để giảm số lượng mail tồn trong queue?</p>										
<p>Trong nguyên lý CAP của distribute system, Redis thuộc nhóm nào sau đây?</p>										
<p>Hệ thống của bạn chỉ có 1 API có kết nối đến MySQL để xử lý dữ liệu, sau khi deploy một tính năng mới, phía application xảy ra lỗi “Too many connections”. Bạn sẽ có thể thực hiện những biện pháp nào để fix vấn đề?</p>										
<p>Khi server (tầng application) của 1 services bị hiện tượng OOM (Out of memory) nguyên nhân có thể do đâu? (Chọn nhiều đáp án)</p>										
<p>Đâu là những thước đo hữu ích cho Performance Testing?</p>										
<p>Phát biểu nào sau đây là đúng khi nói về việc thêm header "Connection: keep-alive	trong	HTTP?	(Chọn	nhiều	đáp	án)</p>"				
<p>Câu nào sau đây mô tả sai về cookie và session?</p>										
<p>Trong domain http://www.google.com thành phần nào được gọi là protocol?</p>										
<p>CCU là chỉ số gì khi đo lường performance của 1 hệ thống?</p>										
<p>Chỉ số nào dưới đây được đo bằng MTBF (Mean Time Netween Failures)?</p>										
<p>Đâu không phải là một tiêu chí đo lường performance của 1 services?</p>										
<p>Hệ thống có thể mở rộng để đáp ứng được việc tăng tải được gọi là?</p>										
<p>Phương án rate limit để tránh thất thoát dữ liệu trong lập trình?</p>										
<p>Dữ liệu ở token JWT khi giải mã base64 có thể thấy được các scope, ta có thể sửa scope rồi mã hóa lại token để lấy thêm quyền không? Vì sao?</p>										
<p>Dự án cần lựa chọn phương án để mã hóa giá trị căn cước công dân trong database, phương án mã hóa nào sau đây là hợp lý và an toàn?</p>										
<p>Các phương án nào sau đây có thể làm thay đổi IP public của người gọi request (để giả mạo IP)?</p>										
<p>Lỗ hổng đọc file bất kỳ ở ứng dụng chạy trên linux có thể đọc được các giá trị gì sau đây (ứng dụng không chạy quyền root)?</p>										
<p>Khi xử lý parse file xml cần lưu ý điều gì để tránh lỗ hổng XXE (XML external entity)</p>										
<p>Các biện pháp để tăng cường tính an toàn cho xác thực đa nhân tố là?</p>										
<p>Cho api như sau: GET /my-user-info trả về thông tin của user hiện tại, nhưng khi truyền thêm param ?user_id=123 thì lại thấy trả về thông tin của user 123, đây là lỗ hổng gì?</p>										
<p>Chỉ thị nào sau đây cho phép accept cookie từ trình duyệt trong cross site request</p>										
<p>Các HTTP Header nào sau đây không được tùy biến giá trị bằng javascript trên trình duyệt cài đặt mặc định?</p>										
<p>Thuật toán mã hóa nào dưới đây khác với các thuật toán còn lại?</p>										
<p>Trong trường hợp cần lưu thông tin token của tổ chức tín dụng để thực hiện thanh toán định kì thì thuật toán mã hóa nào là phù hợp?</p>										
<p>Ứng dụng có sử dụng mã hóa với dữ liệu trong database, trong quá trình xử lý và trao đổi dữ liệu thì cần sử dụng key để mã hóa/giải mã, phương án lưu trữ key nào trong các phương án trên là hợp lý nhất?</p>										
<p>Phương án nào để chống hoặc giảm thiểu tấn công vét cạn?</p>										
<p>Phương pháp nào sau đây dùng để chống lỗ hổng XSS (Cross Site Scripting) hợp lý nhất?</p>										
<p>Redis sử dụng cú pháp nào để xóa toàn bộ keys trong 1 DB?</p>										
<p>Sử dụng câu lệnh nào sau đây để cập nhật cấu hình của redis?</p>										
<p>Redis hỗ trợ bao nhiêu loại dữ liệu?</p>										
<p>Giá trị chuỗi có thể lưu trong redis có kích thước tối đa bao nhiêu?</p>										
<p>Số lượng replication factor tối đa cho 1 topic trong 1 cụm có 9 brokers là?</p>										
<p>Câu lệnh để thực hiện tạo 1 topic trên kafka là?</p>										
<p>Đặc trưng của Kafka là?</p>										
<p>Cấu hình nào chỉ định thời gian giữ logs trong segment trước khi bị xóa đi trong kafka?</p>
<p>Communication port mặc định của Elasticsearch là?</p>
<p>Inverted Index được xây dựng và cập nhật bằng cách sử dụng phân tách các word trong documents? Quá trình phân tách các từ gọi là gì?</p>
<p>Ý nghĩa fuzzy query trong Elasticsearch là gì?</p>										
<p>NRT search trong Elasticsearch là gì?</p>										
<p>Trong ELK, ý nghĩa chữ L là gì?</p>										
<p>Kỹ thuật checksum trong Git sử dụng là gì?</p>										
<p>Công cụ nào sau đây để thực hiện automation testing trên ứng dụng web?</p>										
<p>Cú pháp git checkout dùng để làm gì?</p>										
<p>Continuous Integration gồm những bước nào?</p>										
<p>Mở rộng hệ thống bằng cách bổ sung thêm nhiều nodes được gọi là?</p>										
<p>Trong nguyên lý CAP của distribute system, chữ C thể hiện cho?</p>										
<p>HTTP status nào sau đây là sai?</p>										
<p>HTTP method nào được sử dụng để tạo 1 mới resource?</p>										
<p>Cấu hình tham số URL RESTful để tìm kiếm 1 đơn hàng thông qua order như thế nào?</p>										
<p>Password của users được lưu trữ ở đâu?</p>
<p>Mô hình mạng máy tính nào thường được sử dụng trong tòa nhà?</p>
<p>Trong OSI tầng nào cung cấp dịch vụ cho người dùng?</p>
<p>Kích thước của địa chỉ MAC là?</p>
<p>IP nào sau đây thuộc lớp B?</p>
<p>Facade thuộc nhóm design patterns nào?</p>										
<p>Tính chất nào trong OOP nhằm giảm thiểu việc các class lồng nhau?</p>										
<p>Thành phần private của lớp là?</p>										
<p>Trong lập trình hướng đối tượng các hàm có thể trùng tên nhau gọi là gì?</p>										
<p>Sơ đồ nào sau được biểu diễn hoạt động của hệ thống?</p>										
