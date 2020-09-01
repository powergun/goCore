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

func TestRaiseError(t *testing.T) {
	err := fmt.Errorf("iddqd %s", "e1m1")
	assert.NotEqual(t, nil, err)
}

func TestRaiseException(t *testing.T) {
	assert.NotEqual(t, nil, strstr(""))
}
