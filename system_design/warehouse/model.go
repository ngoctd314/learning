package main

type RelateItem struct {
	Relate      map[uint32]uint64 `bson:"relate"`
	Coefficient []uint64          `bson:"coefficient"`
	ItemID      uint32            `bson:"item_id"`
}
