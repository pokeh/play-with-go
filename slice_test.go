package main

import "testing"

const n = 10

func BenchmarkCap(b *testing.B) {
	sl := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sl = append(sl, i)
	}
}

func BenchmarkLen(b *testing.B) {
	sl := make([]int, n)
	for i := range sl {
		sl[i] = i
	}
}

func BenchmarkNone(b *testing.B) {
	var sl []int
	for i := 0; i < n; i++ {
		sl = append(sl, i)
	}
}
