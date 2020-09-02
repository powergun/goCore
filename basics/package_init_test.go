package basics

// import foobar will trigger the package init function
// which populates the value of ID

import (
	"github.com/powergun/goCore/basics/foobar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageInit(t *testing.T) {
	assert.Equal(t, 101, foobar.ID)
}
