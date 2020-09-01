package sequences

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// byte slice is often used in the IO functions that
// read/write data from/to the outside world

func TestBytesSlice(t *testing.T) {
	f, err := os.Open("/etc/passwd")
	assert.Equal(t, nil, err)

	const sz = 16
	b := make([]byte, sz)

	// Read() is aware of the size of the byte array
	// Read() takes a reference of the byte array
	n, err := f.Read(b)
	assert.Equal(t, nil, f.Close())
	assert.Equal(t, nil, err)
	assert.Equal(t, sz, n)

	// print the raw bytes with nice spacing,
	// not very useful for text file
	//fmt.Printf("% x\n", b)

	// convert byte array into a string (null termination)
	sb := string(b)
	assert.Equal(t, "root:x:0:0:root:", sb)

	// f.Write() only accepts byte slice
	// f.Write([]byte(sb))
}
