package parallel

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewNotification(t *testing.T) {
	no := NewNotification()
	flag := false
	go func() {
		time.Sleep(time.Millisecond * 100)
		flag = true
		no.Done()
	}()
	no.Wait()
	assert.Equal(t, true, flag)
}
