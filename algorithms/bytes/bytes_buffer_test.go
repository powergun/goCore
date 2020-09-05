package bytes

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

// a bytes.Buffer implements the io.Writer interface
// therefore it can be used with other writers in a map
// any interface can be a map key

func TestWriterInterfaceKeys(t *testing.T) {
	output := make(map[io.Writer]bool)
	output[os.Stdout] = false

	b := new(bytes.Buffer)
	output[b] = true

	for w, enabled := range output {
		if enabled {
			if _, err := fmt.Fprintf(w, "iddqd"); err != nil {
				t.Error(err)
			}
		}
	}
	// 5 bytes have been written to the buffer
	// stdout is disabled, so no output shows up in the console
	assert.Equal(t, 5, b.Len())
}
