package benchmark

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	// "warehouse/vendors/roaring"
	"github.com/RoaringBitmap/roaring"
)

// bench Reset on roaring

func Test_resetRoaring(t *testing.T) {
	var data []byte
	//
	{
		ti := time.Now()
		n := 1000
		listRb := make([]*roaring.Bitmap, n)
		for i := 0; i < n; i++ {
			rb := roaring.New()
			relateItemList := make([]uint32, 5000)
			for k := 0; k < 5000; k++ {
				relateItemList[k] = uint32(rand.Intn(1e8) + 1)
			}
			rb.AddMany(relateItemList)
			listRb[i] = rb
		}

		roaring.ParOr(3, listRb...)
		fmt.Println("since: ", time.Since(ti))
	}

	{
		ti := time.Now()
		n := 1000
		listRb := make([]*roaring.Bitmap, n)
		for i := 0; i < n; i++ {
			rb := roaring.New()
			relateItemList := make([]uint32, 5000)
			for k := 0; k < 5000; k++ {
				relateItemList[k] = uint32(rand.Intn(1e7) + 1)
			}
			rb.AddMany(relateItemList)
			listRb[i] = rb
		}
		roaring.ParOr(3, listRb...)

		fmt.Println("since: ", time.Since(ti))
	}
	rb := roaring.New()
	relateItemList := make([]uint32, 5000)
	for k := 0; k < 5000; k++ {
		relateItemList[k] = uint32(rand.Intn(1e8) + 1)
	}
	rb.AddMany(relateItemList)
	fmt.Println(rb.GetSizeInBytes())

	rb1 := roaring.New()
	relateItemList = make([]uint32, 500)
	for k := 0; k < 500; k++ {
		relateItemList[k] = uint32(rand.Intn(1e7) + 1)
	}
	rb1.AddMany(relateItemList)
	fmt.Println(rb1.GetSizeInBytes())

	// rb.Clear()
	// fmt.Println(len(data))
	//
	// pool := sync.Pool{
	// 	New: func() any {
	// 		return roaring.NewWithSize()
	// 	},
	// }
	// l := make([]*roaring.Bitmap, 10000)
	// for i := 0; i < 10000; i++ {
	// 	// pool.Put(roaring.NewWithSize())
	// 	l[i] = roaring.New()
	// 	l[i].FromBuffer(data)
	// 	l[i].Clear()
	// }
	//
	// b.ResetTimer()

	ti := time.Now()
	for i := 0; i < 1; i++ {
		newrb := roaring.New()
		_, err := newrb.FromBuffer(data)
		if err != nil {
			log.Printf("error occur when ReadFrom bitmap (%v)", err)
			return
		}
	}
	fmt.Printf("since: %v ms\n", time.Since(ti).Milliseconds())
}

// func Benchmark_withoutResetRoaring(b *testing.B) {
// 	var data []byte
//
// 	rb := roaring.New()
// 	relateItemList := make([]uint32, 5000)
// 	for k := 0; k < 5000; k++ {
// 		relateItemList[k] = uint32(rand.Intn(1e8) + 1)
// 	}
// 	rb.AddMany(relateItemList)
// 	tmp, err := rb.ToBytes()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	data = tmp
// 	rb.Clear()
//
// 	b.ResetTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		poolRb := roaring.New()
// 		poolRb.FromBuffer(data)
// 	}
// }
