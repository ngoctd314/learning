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

	// insert(ctx, coll)
	query(ctx, coll)
}

func insert(ctx context.Context, coll *mongo.Collection) {
	docs := []interface{}{
		bson.D{
			{"_id", 1},
			{"item", nil},
		},
		bson.D{
			{"_id", 2},
		},
	}
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Println(err)
	}
}

// if you do not specify a projection document, the db.collection.find() method returns all fields
// in the matching documents
func query(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.D{{"item", bson.D{{"$exists", false}}}})
	if err != nil {
		log.Println(err)
	}
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
