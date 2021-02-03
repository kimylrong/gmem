package gmem

import (
	"testing"
)

func BenchmarkMallocWithSize16B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 16)
		Free(buf)
	}
}

func BenchmarkMallocWithSize1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024)
		Free(buf)
	}
}

func BenchmarkMallocWithSize4KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024*4)
		Free(buf)
	}
}

func BenchmarkMallocWithSize1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024*1024)
		Free(buf)
	}
}

func BenchmarkCacheWithSize16B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 16)
		Free(buf)
	}
}

func BenchmarkCacheWithSize1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024)
		Free(buf)
	}
}

func BenchmarkCacheWithSize4KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, _ := MallocWithSize(0, 1024*4)
		Free(buf)
	}
}
