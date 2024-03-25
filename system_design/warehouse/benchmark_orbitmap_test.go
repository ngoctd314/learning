package main

import (
	"math/rand"
	"testing"

	"github.com/RoaringBitmap/roaring"
)

// func Test_timeOrBit(t *testing.T) {
// 	items := 5000
// 	relates := 5000
//
// 	list := make([]*roaring.Bitmap, 0, items*relates)
// 	m := make(map[uint32]struct{})
// 	for i := 0; i < items; i++ {
// 		rb := roaring.New()
// 		tmp := make([]uint32, 0, relates)
// 		for j := 0; j < relates; j++ {
// 			n := rand.Intn(2e7)
// 			tmp = append(tmp, uint32(n))
// 			m[uint32(n)] = struct{}{}
// 		}
// 		rb.AddMany(tmp)
// 		list = append(list, rb)
// 	}
//
// 	now := time.Now()
// 	defer func() {
// 		log.Println("since: ", time.Since(now))
// 	}()
// 	roaring.ParOr(20, list...)
//
// 	// wg := sync.WaitGroup{}
// 	// wg.Add(worker)
// 	// for i := 0; i < worker; i++ {
// 	// 	go func(i int) {
// 	// 		defer wg.Done()
// 	// 		roaring.ParOr(5, listW[i]...)
// 	// 	}(i)
// 	// }
// 	// wg.Wait()
// }

// func Test_maxBytesForBitMap(t *testing.T) {
// 	rb := roaring.New()
// 	list := make([]uint32, 5000)
// 	for i := 0; i < 5000; i++ {
// 		list[i] = uint32(i)
// 	}
// 	rb.AddMany(list)
// 	b, err := rb.ToBytes()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(len(b))
// 	// fmt.Println(rb.String())
// }

func Benchmark_orbit(b *testing.B) {
	items := 1000
	list := make([]*roaring.Bitmap, items)
	for k := 0; k < items; k++ {
		tmp := roaring.New()
		list1 := make([]uint32, 5000)
		for i := 0; i < 5000; i++ {
			list1[i] = uint32(rand.Intn(1e8))
		}
		tmp.AddMany(list1)
		list[k] = tmp
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		roaring.ParOr(6, list...)
	}
}

func Benchmark_add2(b *testing.B) {
	x, y := 100, 1000
	var z int
	for i := 0; i < b.N; i++ {
		for j := 0; j < 375000; j++ {
			x = x + y
			_ = z
		}
	}
}
