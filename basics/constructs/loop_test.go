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

func TestCStyleForLoop(t *testing.T) {
	sum := 0
	for idx := 0; idx < 10; idx++ {
		sum += idx
	}
	assert.Equal(t, sum, 45)

	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		assert.Less(t, i, j)
	}
}

func TestInfiniteLoop(t *testing.T) {
	var counter int
	for {
		counter += 1
		if counter > 10 {
			break
		}
	}
}
