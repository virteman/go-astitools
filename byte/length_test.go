package astibyte_test

import (
	"testing"

	"github.com/virteman/go-astitools/byte"
	"github.com/stretchr/testify/assert"
)

func TestToLength(t *testing.T) {
	assert.Equal(t, []byte("test"), astibyte.ToLength([]byte("test"), ' ', 4))
	assert.Equal(t, []byte("test"), astibyte.ToLength([]byte("testtest"), ' ', 4))
	assert.Equal(t, []byte("test  "), astibyte.ToLength([]byte("test"), ' ', 6))
	assert.Equal(t, []byte("      "), astibyte.ToLength([]byte{}, ' ', 6))
}
