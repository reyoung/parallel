package parallel

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewFuture(t *testing.T) {
	complete := false
	fut := NewFuture(func() (i interface{}, err error) {
		time.Sleep(100 * time.Millisecond)
		complete = true
		return nil, nil
	})
	assert.Equal(t, false, fut.IsComplete())
	_, err := fut.Get()
	assert.Nil(t, err)
	assert.Equal(t, true, complete)
	assert.Equal(t, true, fut.IsComplete())
}
