package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	coll := coll(ctx)

	// insertASingleDocument(ctx, coll)
	// insertMultipleDocuments(ctx, coll)
	// selectAllDocumentInACollection(ctx, coll)
	// selectEqualityCond(ctx, coll)
	// selectEqualityCond(ctx, coll)
	// specifyAndCondition(ctx, coll)
	// specifyORCondition(ctx, coll)
	// insert(ctx, coll)
	// matchAnArray(ctx, coll)
	// matchAnArrayAll(ctx, coll)
	// queryAllDocumentsContainsKey(ctx, coll)
	useQueryOperatorOnArray(ctx, coll)
}

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

func matchAnEmbeddedNestedDocument(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.D{
		{"size", bson.D{
			{"h", 14},
			{"w", 21},
			{"uom", "cm"},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	rs := []any{}
	cur.All(ctx, &rs)
	log.Println(rs)
}

func insert(ctx context.Context, coll *mongo.Collection) {
	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"tags", bson.A{"blank", "red"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"tags", bson.A{"red", "blank"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"tags", bson.A{"red", "blank", "plain"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"tags", bson.A{"blank", "red"}},
			{"dim_cm", bson.A{22.85, 30}},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"tags", bson.A{"blue"}},
			{"dim_cm", bson.A{10, 15.25}},
		},
	}

	result, err := coll.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.InsertedIDs...)
}

func matchAnArray(ctx context.Context, coll *mongo.Collection) {
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
}

func matchAnArrayAll(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.D{{"tags", bson.D{{"$all", bson.A{"red", "blank"}}}}})
	if err != nil {
		log.Fatal(err)
	}
	rs := []any{}
	err = cur.All(ctx, &rs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rs)
}

func queryAllDocumentsContainsKey(ctx context.Context, coll *mongo.Collection) {
	cur, _ := coll.Find(ctx, bson.D{
		{"tags", "red"}, // like {"tags", bson.D{{"$all", bson.A{"red"}}}}
	})
	rs := []any{}
	cur.All(ctx, &rs)
	log.Println(rs)
}

func useQueryOperatorOnArray(ctx context.Context, coll *mongo.Collection) {
	cur, _ := coll.Find(ctx, bson.D{{
		"dim_cm", bson.D{{"$gt", 25}},
	}})
	rs := []any{}
	cur.All(ctx, &rs)
	log.Println(rs)
}

func coll(ctx context.Context) *mongo.Collection {
	clientOpt := options.Client()
	clientOpt.ApplyURI("mongodb://192.168.49.2:30100")

	c, err := mongo.NewClient(clientOpt)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	return c.Database("learn").Collection("crud")
}
