package gmem

import (
	"testing"
)

func BenchmarkMallocWithSize1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024)
		Free(buf)
	}
}

func BenchmarkMallocWithSize128KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024*128)
		Free(buf)
	}
}

func BenchmarkMallocWithSize1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024*1024)
		Free(buf)
	}
}
