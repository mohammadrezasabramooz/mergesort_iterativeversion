package main

import (
	"testing"

	"math/rand"
)

func BenchmarkMergesortiterative_cpu(b *testing.B) {
arr:=make([]int,100000000)
	for i:=0;i<100000000;i++ {
		arr[i]=rand.Int()
	}

	mergesort_iterative_cpu(arr,100000000)


}
