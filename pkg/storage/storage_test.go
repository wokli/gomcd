package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {

	st := CreateStorage()
	st.Add("a", "ololo")
	val, ok := st.Get("a")
	assert.Equal(t, val, "ololo")
	assert.True(t, ok)
}

func TestProcessCommand(t *testing.T) {
	st := CreateStorage()

	res, ok := ProcessCommand(st, "SET a ololo")
	assert.Nil(t, res)
	assert.True(t, ok)

	res, ok = ProcessCommand(st, "GET a")
	assert.True(t, ok)
	assert.Equal(t, *res, "ololo")

	res, ok = ProcessCommand(st, "DEL a")
	assert.Nil(t, res)
	assert.True(t, ok)

	res, ok = ProcessCommand(st, "GET a")
	assert.False(t, ok)
	assert.Nil(t, res)
}
