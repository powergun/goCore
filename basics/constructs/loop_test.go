package logic

import (
	"github.com/smartystreets/gunit"
	"testing"
)

type ForLoopFixture struct {
	*gunit.Fixture
}

func (this *ForLoopFixture) TestRangeBasedFor() {
	nums := [3]int{1, 2, 3}
	for i := range nums {
		i += i
	}
}

func (this *ForLoopFixture) TestIndexBasedFor() {
	for i:=1; i < 10; i++ {
		i += 1
	}
}

func TestForLoopFixture(t *testing.T) {
	gunit.Run(new(ForLoopFixture), t)
}
