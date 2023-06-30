package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	coll := coll(ctx)

	// delete(ctx, coll)
	// insert(ctx, coll)
	query(ctx, coll)
}

func delete(ctx context.Context, coll *mongo.Collection) {
	coll.DeleteMany(ctx, bson.M{})
}

func insert(ctx context.Context, coll *mongo.Collection) {
	docs := []interface{}{
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "mat"},
			{"qty", 85},
			{"size", bson.D{
				{"h", 27.9},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "mousepad"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 19},
				{"w", 22.85},
				{"uom", "in"},
			}},
			{"status", "P"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "P"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "sketchbook"},
			{"qty", 80},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "sketch pad"},
			{"qty", 95},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}

	rs, err := coll.InsertMany(context.TODO(), docs)
	fmt.Println(rs, err)
}

// if you do not specify a projection document, the db.collection.find() method returns all fields
// in the matching documents
func query(ctx context.Context, coll *mongo.Collection) {
	cur, err := coll.Find(ctx, bson.D{{}})
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
