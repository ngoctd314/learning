package main

func main() {
	for i := 0; i < 5; i++ {
	}
}

// func printAlloc() {
// var m runtime.MemStats
// ReadMemStats populates m with memory allocator statistic
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in constrast with a heap profile
// which is a snapshot as of the most recently completed garbage
// collection cycle.
// runtime.ReadMemStats(&m)
// fmt.Printf("%d KB\n", m.Alloc/1024)
// }
