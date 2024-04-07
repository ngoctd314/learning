1 2 3 4 6 8 10
2 1 3 5
3 1 2
4 2 5
5 1 4


A: 10001
B: 11001

deletedBitmap: 00001

Result = (A|B) ANDNOT deletedBitmap = 11001 ANDNOT 00001 = 11000

+ Ví dụ có item1, item2 

item1: {2, 3, 4} => bit(item1) = 011100

item2: {4, 5}    => bit(item2) = 110000


countBit1(item1 | item2) =countBit1(111100) = 4

+ delete item2 khỏi hệ thống

DELETE item2 FROM relate

Cập nhật deletedBitmap = 000100

+ countBit1(item1 ANDNOT deletedBitmap) = countBit1(11000) = 2
