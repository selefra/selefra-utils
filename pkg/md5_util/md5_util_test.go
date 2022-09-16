package md5_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5String(t *testing.T) {
	s, err := Md5String("Tom")
	assert.Nil(t, err)
	assert.Equal(t, "d9ffaca46d5990ec39501bcdf22ee7a1", s)
}
