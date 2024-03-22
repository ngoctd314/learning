# MongoDB and Document Model

id, item_id, relate, mod (max 1562500)

select by item_id: avg rows: 5000 thuộc (1 -> 1562500)

1562500 / 64  / 64 / 64 ~ 6

=> hash 5000 => bé hơn

lấy 10K item_id => 10000*5000 = 50_000_000

+ query: item1, item2, ...item_n thuộc những tập hash1 chung nào (0 ->  1562500)

item_id, hash1, hash2, hash3

hash1: hash 1562500 => 64 tập nhỏ hơn

id, item_id, hash1
x   1        0
x   1        1
x   1        2

x   2        0
x   2        1
x   2        2

hash1 in (0, 24415)

+ query: item1, item2, ...item_n thuộc những tập hash2 chung nào (0 ->  24415)

hash2: hash1 24415 => 64 tập nhỏ hơn

id, hash1_id, hash2, item_id
x   1         0      
x   1         1
x   1         2

x   2         0
x   2         1
x   2         2

hash2 in (0, 382)

382*10000 = 3820000 op

item_id => hash2 => hash1 (tim vị trí = 1)

Ví dụ

id, item_id, hash1
1   1        0
1   4        0

SELECT item_id, hash1_id FROM hash2 where item_id in (1, 4) GROUP By hash2 HAVING count(item_id) > 0

item_id, hash2, hash1_id
1        0      1
4        0      2

SELECT item_id, hash1_id FROM hash1 where id in (1, 2) GROUP By hash1 HAVING count(item_id) > 0
