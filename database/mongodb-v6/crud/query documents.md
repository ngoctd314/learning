# Query Documents

## 1. Overview

### 1.1. Select All Documents in a Collection

To select all documents in the collection, pass an empty document as the query filter parameter to the find method.

```go
func selectAllDocumentInACollection(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	rs := []any{}
	if err := cur.All(ctx, &rs); err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}
```

### 1.2. Specify Equality Condition

```go
func selectEqualityCond(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.M{"status": "D"})
	if err != nil {
		log.Fatal(err)
	}

	rs := []any{}
	if err := cur.All(ctx, &rs); err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}
```

### 1.3. Specify Conditions Using Query Operators

A query filter document can use the query operators to specify conditions

```go
func queryWithCond(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.M{"status": bson.M{"$in": []any{"A", "D"}}})
	if err != nil {
		log.Fatal(err)
	}

	rs := []any{}
	if err := cur.All(ctx, &rs); err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}

func specifyAndCondition(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.M{"status": "A", "qty": bson.M{"$lt": 30}})
	if err != nil {
		log.Fatal(err)
	}

	rs := []any{}
	if err := cur.All(ctx, &rs); err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}

func specifyORCondition(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.M{"$or": []any{
		bson.M{"status": "A"},
		bson.M{"qty": bson.M{"$lt": 30}},
	}})
	if err != nil {
		log.Fatal(err)
	}

	rs := []any{}
	if err := cur.All(ctx, &rs); err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}
```

### 1.4. Behavior

**Cursor**

The db.collection.find() method returns a cursor to the matching documents

**Read Isolation**

For reads to replica sets and replica set shards, read concern allow clients to choose a level of isolation for their reads.

## 2. Query on Embedded/Nested Documents

## 3. Query an array

### 3.1. Match an array

To specify equality condition on an array, use the query document {<field>: <value>} where <value> is the exact array to match, including the order of the elements.

The following example queries for all documents where the field tags value is an array with exactly two elements "red" and "blank"

```go
cur, err := coll.Find(ctx, bson.D{{"tags", bson.A{"red", "blank"}}})
if err != nil {
	log.Fatal(err)
}
rs := []any{}
err = cur.All(ctx, &rs)
if err != nil {
	log.Fatal(err)
}
log.Println(rs)
```

If, instead you wish to find an array that contains both the element "red" and "blank" without regard to order or other elements in the array, use the $all operator

```go
cursor, err := coll.Find(
	context.TODO(),
	bson.D{
		{"tags": bson.D{{"$all", bson.A{"red", "blank"}}}}
	})
```
### 3.2. Query an Array of an Element

To query if the array field contains at least one element with the specified value, use the filter `{<field>: <value>}` where `<value>` is the element value.

```go
cur, _ := coll.Find(ctx, bson.D{
	{"tags", "red"}, // like {"tags", bson.D{{"$all", bson.A{"red"}}}}
})
rs := []any{}
cur.All(ctx, &rs)
log.Println(rs)
```

To specify conditions on the elements in the array field, use query operations in the query filter documents

For example, the following operation queries for all documents where the array dim_cm contains at least one element whose value is greater than 25.

```go
cur, _ := coll.Find(ctx, bson.D{{
	"dim_cm", bson.D{{"$gt", 25}},
}})
rs := []any{}
cur.All(ctx, &rs)
log.Println(rs)
```

### 3.3. Specify Multiple Conditions for Array Elements

You can specify the query such that either a single array element meets these condition or any combination of array elements meets the conditions.

**Query an Array with Compound Filter Conditions on the Array Elements**

One element can satisfy the greater than 15 condition and another element can satisfy the less than 20 condition, or a single element can satisfy both.

```go
cur, err := coll.Find(ctx, bson.D{
	{"dim_cm", bson.D{{"$gt", 15}, {"$lt", 20}}}
})
```

**Query for an Array element that meets multiple criteria**

Use `$elemMatch` operator to specify multiple criteria on the element of an array that at least one array element satisfies all the specified criteria.

```go
cur, err := coll.Find(ctx, bson.D{
	{"dim_cm", bson.D{
		{"$elemMatch", bson.D{
			{}
		}}
	}}
})
```

**Query for an Element by the Array index position**

**Query an Array by Array length**