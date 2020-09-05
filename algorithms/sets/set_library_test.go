package sets

// https://github.com/scylladb/go-set
import (
	"github.com/scylladb/go-set/strset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetLibrary(t *testing.T) {
	s1 := strset.New("entry 1", "entry 2")
	s2 := strset.New("entry 2", "entry 3")
	s3 := strset.Intersection(s1, s2)
	assert.False(t, s3.Has("entry 1"))
	assert.True(t, s3.Has("entry 2"))
}
