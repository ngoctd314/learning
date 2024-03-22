package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoConn(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.49.2:30302"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	return client
}

func cleanDB(client *mongo.Client) {}

func main() {
	ctx := context.Background()

	client := initMongoConn(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
	relateItemCollection := client.Database("warehouse").Collection("relate_item")

	// genData()
	// insertData(relateItemCollection)
	// return

	r := relateItemRepository{collection: relateItemCollection}

	// for i := 1; i <= 5; i++ {
	// 	if err := r.createItem(ctx, uint32(i)); err != nil {
	// 		log.Printf("insertRelate error (%v)\n", err)
	// 	}
	// }
	rs := r.countDistinct(ctx, 1, 2)
	fmt.Println(rs)
}
