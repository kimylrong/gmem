package gmem

import (
	"fmt"
	"sync"
)

const (
	maxSizeBucket1 = 1024               //1KB
	maxSizeBucket2 = 32 * 1024          //32KB
	maxSizeBucket3 = 1024 * 1024        //1MB
	maxSizeBucket4 = 32 * 1024 * 1024   //32MB
	maxSizeBucket5 = 1024 * 1024 * 1024 //1GB

	alignmentBucket1 = 8               //8B
	alignmentBucket2 = 256             //256B
	alignmentBucket3 = 8 * 1024        //8KB
	alignmentBucket4 = 256 * 1024      //256KB
	alignmentBucket5 = 8 * 1024 * 1024 //8MB
)

type bucket struct {
	level     int
	min       int
	max       int
	alignment int
	pools     []*sync.Pool
}

var (
	errCapNotValid  = fmt.Errorf("cap should between 1 and %d", maxSizeBucket5)
	errSizeNotValid = fmt.Errorf("size not valid")
)

var buckets [5]*bucket

func init() {
	buckets[0] = newBucket(1, 0, maxSizeBucket1, alignmentBucket1)
	buckets[1] = newBucket(2, maxSizeBucket1, maxSizeBucket2, alignmentBucket2)
	buckets[2] = newBucket(3, maxSizeBucket2, maxSizeBucket3, alignmentBucket3)
	buckets[3] = newBucket(4, maxSizeBucket3, maxSizeBucket4, alignmentBucket4)
	buckets[4] = newBucket(5, maxSizeBucket4, maxSizeBucket5, alignmentBucket5)
}

func newBucket(level, min, max, alignment int) *bucket {
	b := &bucket{
		level:     level,
		min:       min,
		max:       max,
		alignment: alignment,
	}
	l := b.poolSize()
	b.pools = make([]*sync.Pool, l, l)
	for i := 0; i < l; i++ {
		j := i
		b.pools[i] = &sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, b.min+b.alignment*(j+1))
			},
		}
	}
	return b
}

func (c *bucket) poolSize() int {
	return (c.max - c.min) / c.alignment
}

func (c *bucket) malloc(cap int) []byte {
	p := c.pools[c.index(cap)]
	return p.Get().([]byte)
}

func (c *bucket) free(buf []byte) bool {
	cap := cap(buf)
	if (cap-c.min)%c.alignment != 0 {
		// not alignment, just ignore
		return false
	}
	buf = buf[:0]
	p := c.pools[c.index(cap)]
	p.Put(buf)
	return true
}

func (c *bucket) index(cap int) int {
	return (cap - c.min - 1) / c.alignment
}

func level(cap int) int {
	if cap <= 0 || cap > maxSizeBucket5 {
		return -1
	}
	if cap <= maxSizeBucket1 {
		return 1
	} else if cap <= maxSizeBucket2 {
		return 2
	} else if cap <= maxSizeBucket3 {
		return 3
	} else if cap <= maxSizeBucket4 {
		return 4
	} else if cap <= maxSizeBucket5 {
		return 5
	}
	return -1
}

// Malloc fetch a []byte from pool with cap. size been set to 0.
func Malloc(cap int) ([]byte, error) {
	return MallocWithSize(0, cap)
}

// MallocWithSize fetch a []byte from pool with size and cap.
func MallocWithSize(size, cap int) ([]byte, error) {
	level := level(cap)
	if level == -1 {
		return nil, errCapNotValid
	}
	if size < 0 || size > cap {
		return nil, errSizeNotValid
	}
	c := buckets[level-1]
	buf := c.malloc(cap)
	buf = buf[:size]
	return buf, nil
}

// Free release []byte back to pool
func Free(buf []byte) bool {
	level := level(cap(buf))
	if level == -1 {
		// ignore
		return false
	}
	c := buckets[level-1]
	return c.free(buf)
}
