package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.RemoveAll("./workdir")
		os.Mkdir("./workdir", os.ModeDir)
		start := time.Now()
		Concurrency()
		end := time.Now()
		fmt.Printf("Time elapsed for concurrency: %v \n", int(end.Sub(start).Seconds()))
	}
}
func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.RemoveAll("./workdir")
		os.Mkdir("./workdir", os.ModeDir)
		start := time.Now()
		Sequential()
		end := time.Now()
		fmt.Printf("Time elapsed for sequential: %v \n", int(end.Sub(start).Seconds()))
	}
}
