package logic

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConstruct(t *testing.T) {
	title := struct {
		English string
		French  string
	}{"IDDQD", "IDDQD"}
	assert.Equal(t, "IDDQD", title.English)
}

// use anonymous interface for runtime reflection
func TestReflection(t *testing.T) {
	factory := func() interface{} {
		b := strings.Builder{}
		return b
	}
	b := factory()
	if _, ok := b.(interface {
		String() string
	}); !ok {
	}
}
