package basics

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	EmptyString = errors.New("empty string")
)

func strstr(s string) error {
	return EmptyString
}

func TestRaiseArbitraryError(t *testing.T) {
	err := fmt.Errorf("iddqd %s", "e1m1")
	assert.NotEqual(t, nil, err)
}

func TestRaiseDefinedError(t *testing.T) {
	assert.NotEqual(t, nil, strstr(""))

	switch strstr("") {
	case EmptyString:
		assert.Equal(t, 0, 0)
	}
}
