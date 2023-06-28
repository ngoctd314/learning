# Quey on Embedded/Nested Documents

<<<<<<< HEAD
## Match an Embedded/Nested Document

To specify an equality condition on a field that is an embedded/nested document, use the query filter document  {<field>: <value>} when value is the document to match.

=======
## Match an Embedded/Nested Document (full document match)

Use the query filter document {<field>: <value>} where <value> is the document to match

```go
db.inventory.find({size: {h: 14, w: 21, uom: "cm"}})
```

Equality matches on the whole embedded document require an exact match of the specified document, including the field order. The following query does not match any documents

```go
do.inventory.find({size: {w: 21, h: 14, uom: "cm"}})
```

## 
>>>>>>> 5b6d0c495e7a98b75adf8449125772bf64d8e378
