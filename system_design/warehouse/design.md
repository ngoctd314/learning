- Database: MongoDB

- Schema:


related_item

insert(relateID): insert(65), insert(641)

```json
{
    "_id": 1,
    "relate_1": [{x: 1, y: 1}, {x: 10, y: 1}],
    "relate_2": [{x: 1, y: 1}, {x: 11, y: 1}]
}
```

```go
if r1.x == r2.x {
    rs += cnt(r1.y | r2.y)
} else if r1.x < r2.x {
   rs += cnt(r1.x) 
   r1P++
} else {
   rs += cnt(r2.x) 
   r2P++
}
```

Optimize phase 1.

```json
{
    "_id": 1,
    "relate_hash1": [1,1],
    "coefficient_hash1": [2^1+2^10]
}
```

range relate: [1, max_item_id], length: ~5000

range coefficient_hash: [0, max_item_id/64], length: ~5000 (max_item_id: 100M -> range [0, 1M6])

+ max length(relatei), i in [1, 5] = 5000

relatei[j] in [1, 5M] 

hash_i[j] in [1, ...,78124]
