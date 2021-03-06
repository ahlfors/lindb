package strutil

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_GetStringValue(t *testing.T) {
	assert.Equal(t, "sum", GetStringValue("sum"))
	assert.Equal(t, "sum", GetStringValue("'sum'"))
	assert.Equal(t, "'sum", GetStringValue("'sum"))
	assert.Equal(t, "sum", GetStringValue("\"sum\""))
	assert.Equal(t, "", GetStringValue(""))
}

func Test_ByteSlice2String(t *testing.T) {
	s := []byte("abc")
	assert.Equal(t, "abc", ByteSlice2String(s))
}

func Test_String2ByteSlice(t *testing.T) {
	s := "abc"
	assert.Equal(t, []byte("abc"), String2ByteSlice(s))
}
