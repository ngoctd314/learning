package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type relateItemRepository struct {
	collection *mongo.Collection
}

// func (r *relateItemRepository) createItem(ctx context.Context, itemID uint32) error {
// 	document := RelateItem{
// 		Relate:      make(map[uint32]uint64),
// 		Coefficient: make([]uint64, 25),
// 		ItemID:      itemID,
// 	}
//
// 	rs, err := r.collection.InsertOne(ctx, document)
// 	if err != nil {
// 		return err
// 	}
// 	_ = rs
//
// 	return nil
// }

func (r *relateItemRepository) insertRelate(ctx context.Context, itemID uint32, relateIDs ...uint32) error {
	relates := make(map[uint32]uint64)
	co := make([]uint64, chunk+1)
	for _, relateID := range relateIDs {
		relateValue, relatePos := byte(relateID%chunkLen), relateID/chunkLen
		hashValue, hashPos := byte(relatePos%chunkLen), relatePos/chunkLen

		relates[relatePos] |= pow2(relateValue)
		co[hashPos] |= pow2(hashValue)

	}
	document := RelateItem{
		Relate: relates,
		// Coefficient: co,
		ItemID: itemID,
	}

	rs, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	_ = rs

	return nil
}

func (r *relateItemRepository) updateRelate(ctx context.Context, itemID uint32, relateID uint32) error {
	relateValue, relatePos := byte(relateID%chunkLen), relateID/chunkLen
	hashValue, hashPos := byte(relatePos%chunkLen), relatePos/chunkLen

	// insert or update by itemID
	// if create => relateHash need to update, coefficientHash need to update
	// coefficientHash only need to $bit or
	// relateHash need to insert push at pos
	filter := bson.D{
		{Key: "item_id", Value: itemID},
	}
	update := bson.D{
		{
			Key: "$bit",
			Value: bson.D{
				{
					Key:   fmt.Sprintf("relate.%d", relatePos),
					Value: bson.D{{Key: "or", Value: pow2(relateValue)}},
				},
				{
					Key:   fmt.Sprintf("coefficient.%d", hashPos),
					Value: bson.D{{Key: "or", Value: pow2(hashValue)}},
				},
			},
		},
	}
	opts := options.Update().SetUpsert(true)
	rs, err := r.collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return fmt.Errorf("%w, itemID %d, relateID: %d", err, itemID, relateID)
	}
	_ = rs

	return nil
}

func (r *relateItemRepository) countDistinct(ctx context.Context, itemIDs ...uint32) uint64 {
	filter := bson.D{{
		Key: "item_id",
		Value: bson.D{{
			Key:   "$in",
			Value: itemIDs,
		}},
	}}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Printf("find itemIDs error: (%v)\n", err)
		return 0
	}
	defer cur.Close(ctx)

	var results []RelateItem
	for cur.Next(ctx) {

		var elem RelateItem
		if err := cur.Decode(&elem); err != nil {
			log.Printf("decode itemID error: (%v)\n", err)
			continue
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Printf("countDistinct cur.Err error: (%v)\n", err)
	}

	now := time.Now()
	defer func() {
		log.Println("since: ", time.Since(now).Seconds())
	}()

	// hash range
	var finalResult uint64
	// trackItemID := make(map[uint32]int)
	// chunk
	// wg := sync.WaitGroup{}
	// wg.Add(chunk + 1)
	// ch := make(chan uint64, (chunk+1)*chunkLen)
	m := make(map[uint64]uint64)
	fmt.Println("len results", len(results))
	for i := range results {
		k := results[i].Relate
		for k, v := range k {
			// fmt.Println(n, v)
			m[uint64(k)] |= v
		}
	}
	for _, v := range m {
		finalResult += countBit1(v)
	}
	// for i := 0; i <= chunk; i++ {

	// go func(i int) {
	// 	defer wg.Done()

	// var j byte = 0
	// each chunk from 0 -> 63
	// for j = 0; j < chunkLen; j++ {
	// 	var chunk uint64
	// 	for _, r := range results {
	// 		co := r.Coefficient
	//
	// 		if co[i] == 0 {
	// 			continue
	// 		}
	// 		if !sameAtBit(co[i], pow2(j)) {
	// 			continue
	// 		}
	// 		chunk |= r.Relate[chunkLen*uint32(i)+uint32(j)]
	// 	}
	// 	finalResult += countBit1(chunk)
	// }
	// }(i)
	// }
	// wg.Wait()
	// close(ch)
	// for v := range ch {
	// 	finalResult += v
	// }

	return finalResult
}

func (r *relateItemRepository) drop() {
	r.collection.Drop(context.Background())
}
