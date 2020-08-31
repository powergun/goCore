package logic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitchCase(t *testing.T) {
	num_bytes, err := fmt.Printf("")
	switch {
	case err != nil:
		assert.Equal(t, num_bytes, 0)
	case num_bytes == 0:
		assert.Equal(t, err, nil)
	default:
		fmt.Println("iddqd")
	}
	assert.Equal(t, "", "")
}

func TestSwitchCaseVariable(t *testing.T) {
	for i, r := range "iddqddkfa" {
		switch r {
		case 'i':
			assert.Equal(t, i, 0)
		case 'd', 'q':
			assert.Greater(t, i, 0)
		default:
			assert.Greater(t, i, 3)
		}
	}
}
