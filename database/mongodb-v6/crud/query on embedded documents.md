# Quey on Embedded/Nested Documents

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