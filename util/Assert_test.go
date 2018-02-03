package util

import (
	"testing"
)

func BenchmarkAssert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Assert(2 > 1, "2 > 1")
	}
}

func BenchmarkNoAssert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
