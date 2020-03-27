package parallel_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"parallel"
	"strings"
	"testing"
)

func TestForGood(t *testing.T) {
	N := 10
	result := make([]bool, N)
	for i := 0; i < N; i++ {
		result[i] = false
	}
	assert.Nil(t, parallel.For(10, func(i int, total int) error {
		result[i] = true
		return nil
	}))
}
func TestForBad(t *testing.T) {
	err := parallel.For(10, func(i int, total int) error {
		if i%2 == 0 {
			return nil
		}
		return errors.New("test error")
	})
	assert.NotNil(t, err)
	result := strings.Split(err.Error(), "\n")
	assert.Equal(t, 6, len(result))
}
