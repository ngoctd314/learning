[[_TOC_]]

TLDR; kafka-go is more performance

## Producer

### Sequential

|lib|Time Per Operation (ns/op)|Memory Per Operation (B/op)|Allocations Per Operation(allocs/op)|
|-|-|-|-|
|kafka-go|19613|3211|39|
|sarama|88908|3724|69|

**Using a sequential process `kafka-go` is**

- Operation faster 4.5 times when compared to `sarama`.
- Memory used fewer 0.43 times when compared to `sarama`.
- Memory allocation fewer 0.137 times when compared to `sarama`.


**kafka-go**

```go
func Benchmark_sequentialProduce(b *testing.B) {
	w := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "test",
		BatchSize: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
        w.WriteMessages(context.Background(),
            kafka.Message{
                Value: []byte("test"),
            },
        )
	}
}
```

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/kafka-go

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               71848             19711 ns/op            3200 B/op         39 allocs/op
PASS
ok      lkafka/lib/kafka-go     1.595s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               80811             22364 ns/op            3241 B/op         39 allocs/op
PASS
ok      lkafka/lib/kafka-go     1.968s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               77828             24252 ns/op            3219 B/op         39 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.055s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               89311             17492 ns/op            3194 B/op         39 allocs/op
PASS
ok      lkafka/lib/kafka-go     1.708s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               88671             14247 ns/op            3199 B/op         39 allocs/op
PASS
ok      lkafka/lib/kafka-go     1.410s
```

**sarama**

```go
func Benchmark_sequenceProduce(b *testing.B) {
	w, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
        w.SendMessage(&sarama.ProducerMessage{
            Topic: "test",
            Value: sarama.StringEncoder("test"),
        })
	}
}
```

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/sarama

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               13153             89401 ns/op            3724 B/op         69 allocs/op
PASS
ok      lkafka/lib/sarama       2.107s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               13005             90536 ns/op            3724 B/op         69 allocs/op
PASS
ok      lkafka/lib/sarama       2.074s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               13412             87151 ns/op            3724 B/op         69 allocs/op
PASS
ok      lkafka/lib/sarama       2.084s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               13646             88460 ns/op            3724 B/op         69 allocs/op
PASS
ok      lkafka/lib/sarama       2.107s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequenceProduce
Benchmark_sequenceProduce-12               13585             88992 ns/op            3723 B/op         69 allocs/op
PASS
ok      lkafka/lib/sarama       2.112s
```

### Concurrency

**kafka-go**

```go
func Benchmark_concurrencyProduce(b *testing.B) {
	w := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "test",
		BatchSize: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency := func() {
			level := 6
			wg := sync.WaitGroup{}
			wg.Add(level)
			for c := 0; c < level; c++ {
				go func() {
					defer wg.Done()
					w.WriteMessages(context.Background(),
						kafka.Message{
							Value: []byte("test"),
						},
					)
				}()
			}
			wg.Wait()
		}
		concurrency()
	}
}
```

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/kafka-go

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12            15397            160408 ns/op           21697 B/op        240 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.886s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12            14246            163125 ns/op           22254 B/op        240 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.777s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12            14496            127703 ns/op           22029 B/op        240 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.234s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12            14959            159522 ns/op           21593 B/op        240 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.826s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12            15238            117391 ns/op           21746 B/op        240 allocs/op
PASS
ok      lkafka/lib/kafka-go     2.236s
```

**sarama**

```go
func Benchmark_concurrencyProduce(b *testing.B) {
	w, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency := func() {
			level := 6
			wg := sync.WaitGroup{}
			wg.Add(level)
			for c := 0; c < level; c++ {
				go func() {
					defer wg.Done()
					w.SendMessage(&sarama.ProducerMessage{
						Topic: "test",
						Value: sarama.StringEncoder("test"),
					})
				}()
			}
			wg.Wait()
		}
		concurrency()
	}
}
```

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/sarama

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12             5215            224972 ns/op           12098 B/op        221 allocs/op
PASS
ok      lkafka/lib/sarama       2.179s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12             5602            218694 ns/op           11982 B/op        218 allocs/op
PASS
ok      lkafka/lib/sarama       2.180s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U

Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12             5001            216573 ns/op           11964 B/op        218 allocs/op
PASS
ok      lkafka/lib/sarama       1.115s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12             5505            213545 ns/op           11996 B/op        219 allocs/op
PASS
ok      lkafka/lib/sarama       2.093s

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_concurrencyProduce
Benchmark_concurrencyProduce-12             5588            217933 ns/op           11977 B/op        219 allocs/op
PASS
ok      lkafka/lib/sarama       2.199s
```

### Batch processing

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/kafka-go

```

**sarama**
```go
func Benchmark_sequentialBatch(b *testing.B) {
	cnf := sarama.NewConfig()
	cnf.Producer.Flush.Messages = 100
	cnf.Producer.Flush.Frequency = time.Second
	w, _ := sarama.NewAsyncProducer([]string{"localhost:9092"}, cnf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Input() <- &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder("test"),
		}
	}
}
```

```txt
goos: linux
goarch: amd64
pkg: lkafka/lib/sarama

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
Benchmark_sequentialBatch
Benchmark_sequentialBatch-12                2882            412987 ns/op           65013 B/op        861 allocs/op
PASS
ok      lkafka/lib/sarama       2.171s
```

## Consumer
