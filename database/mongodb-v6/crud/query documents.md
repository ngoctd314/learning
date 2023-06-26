# Query Documents

## Select All Documents in a Collection

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

## Specify Equality Condition

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

## Specify Conditions Using Query Operators

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

## Behavior

**Cursor**

The db.collection.find() method returns a cursor to the matching documents

**Read Isolation**

For reads to replica sets and replica set shards, read concern allow clients to choose a level of isolation for their reads.

