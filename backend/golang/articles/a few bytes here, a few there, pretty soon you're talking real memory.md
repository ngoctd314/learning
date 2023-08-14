# A few bytes here, a few there, pretty soon you're talking real memory

```go
func BenchmarkSortStrings(b *testing.B) {
    s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        sort.Strings(s)
    }
}
```