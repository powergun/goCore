package logic

import (
	"github.com/smartystreets/gunit"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ForLoopFixture struct {
	*gunit.Fixture
}

func (this *ForLoopFixture) TestRangeBasedFor() {
	nums := [3]int{1, 2, 3}
	sum := 0
	// for i := range nums will only loop-through the elements
	for idx, i := range nums {
		sum += i + idx
	}
	assert.Equal(this, 9, sum)
}

func (this *ForLoopFixture) TestIndexBasedFor() {
	for i := 1; i < 10; i++ {
		i += 1
	}
}

func TestForLoopFixture(t *testing.T) {
	gunit.Run(new(ForLoopFixture), t)
}
