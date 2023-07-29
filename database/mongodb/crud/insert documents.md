# Insert Documents

## Insert a Single Document

db.collection.insertOne() inserts a single document into a collection

```go
func insertASingleDocument(ctx context.Context, coll *mongo.Collection) {
	rs, err := coll.InsertOne(ctx, bson.M{
		"item": "canvas",
		"qty":  100,
		"tags": bson.A{"cotton"},
		"size": bson.M{
			"h":   28,
			"w":   35.5,
			"uom": "cm",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	cur, err := coll.Find(ctx, bson.M{"_id": rs.InsertedID})
	if err != nil {
		log.Println(err)
	}
	for cur.Next(ctx) {
		// if the document does not specify an _id field, the driver adds the _id field with an ObjectId value to the new document.
		log.Println(cur.Current.Values())
	}
}

func insertMultipleDocuments(ctx context.Context, coll *mongo.Collection) {
	rs, err := coll.InsertMany(ctx, []interface{}{
		bson.M{
			"item": "journal",
			"qty":  int32(25),
			"tags": bson.A{"blank", "red"},
			"size": bson.M{
				"h":   14,
				"w":   21,
				"uom": "cm",
			},
		},
		bson.M{
			"item": "mat",
			"qty":  int32(25),
			"tags": bson.A{"gray"},
			"size": bson.M{
				"h":   27.9,
				"w":   35.5,
				"uom": "cm",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	cur, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": rs.InsertedIDs}})
	if err != nil {
		log.Fatal(err)
	}
	results := []any{}
	cur.All(ctx, &results)
	log.Println(results...)
}
```

## Insert Behavior

**Collection Creation**

If the collection does not currently exist, insert operations will create the collection.

**_id Field**

In MongoDB, each document stored in a collection requires a unique _id field that acts as a primary key. If an inserted document omits the _id field, the MongoDB driver automatically generates on ObjectId for the _id field.

This is also applies to documents inserted through operations with upsert:true

**Atomicity**

All write operations in MongoDB are atomic on the level of a single document.

**Write Ack**

With write concerns, you can specify the level of ack requested from MongoDB for write operations.