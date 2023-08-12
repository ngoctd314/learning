package main

import (
	"fmt"
	"io"
	"runtime"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %d MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc: %d MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys: %d MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC: %d\n", m.NumGC)
}

func copySourceToDest(source io.Reader, dest io.Writer) error {
	data := make([]byte, 5)
	source.Read(data)

	dest.Write(data)

	return nil
}

type fileReader struct {
	content []byte
}

func (f fileReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = f.content[i]
	}
	return len(f.content), nil
}

type fileWriter struct {
	content []byte
}

func (w *fileWriter) Write(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		w.content = append(w.content, p[i])
	}

	return len(p), nil
}

func main() {
}
