package gmem

import "testing"

func TestMalloc(t *testing.T) {
	cases := []struct {
		cap       int
		expectCap int
	}{
		{
			cap:       1,
			expectCap: 8,
		},
		{
			cap:       7,
			expectCap: 8,
		},
		{
			cap:       8,
			expectCap: 8,
		},
		{
			cap:       9,
			expectCap: 8 + 8,
		},
		{
			cap:       1023,
			expectCap: 1024,
		},
		{
			cap:       1024,
			expectCap: 1024,
		},
		{
			cap:       1025,
			expectCap: 1024 + 256,
		},
		{
			cap:       32*1024 - 1,
			expectCap: 32 * 1024,
		},
		{
			cap:       32 * 1024,
			expectCap: 32 * 1024,
		},
		{
			cap:       32*1024 + 1,
			expectCap: 32*1024 + 8*1024,
		},
		{
			cap:       1024*1024 - 1,
			expectCap: 1024 * 1024,
		},
		{
			cap:       1024 * 1024,
			expectCap: 1024 * 1024,
		},
		{
			cap:       1024*1024 + 1,
			expectCap: 1024*1024 + 256*1024,
		},
	}

	for _, c := range cases {
		buf, _ := Malloc(c.cap)
		if cap(buf) != c.expectCap {
			t.Fail()
		}
	}
}

func TestMallocWithSize(t *testing.T) {
	cases := []struct {
		size       int
		cap        int
		expectSize int
		expectCap  int
	}{
		{
			size:       0,
			cap:        8,
			expectSize: 0,
			expectCap:  8,
		},
		{
			size:       1,
			cap:        8,
			expectSize: 1,
			expectCap:  8,
		},
		{
			size:       0,
			cap:        7,
			expectSize: 0,
			expectCap:  8,
		},
		{
			size:       8,
			cap:        8,
			expectSize: 8,
			expectCap:  8,
		},
		{
			size:       9,
			cap:        9,
			expectSize: 9,
			expectCap:  16,
		},
	}

	for _, c := range cases {
		buf, _ := MallocWithSize(c.size, c.cap)
		if len(buf) != c.expectSize {
			t.Fail()
		}
		if cap(buf) != c.expectCap {
			t.Fail()
		}
	}
}

func TestMallocErr(t *testing.T) {
	cases := []struct {
		size int
		cap  int
	}{
		{
			size: 0,
			cap:  0,
		},
		{
			size: 0,
			cap:  -1,
		},
		{
			size: 0,
			cap:  1024*1024*1024 + 1,
		},
		{
			size: 3,
			cap:  2,
		},
	}

	for _, c := range cases {
		_, err := MallocWithSize(c.size, c.cap)
		if err == nil {
			t.Fail()
		}
	}
}

func TestFree(t *testing.T) {
	cases := []struct {
		buf          []byte
		expectResult bool
	}{
		{
			buf:          make([]byte, 0, 0),
			expectResult: false,
		},
		{
			buf:          make([]byte, 0, 1),
			expectResult: false,
		},
		{
			buf:          make([]byte, 0, 8),
			expectResult: true,
		},
		{
			buf:          make([]byte, 0, 9),
			expectResult: false,
		},
	}

	for _, c := range cases {
		result := Free(c.buf)
		if result != c.expectResult {
			t.Fail()
		}
	}
}
