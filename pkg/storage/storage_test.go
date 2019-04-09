package storage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {

	st := CreateStorage()
	st.Add("a", "ololo")
	val, err := st.Get("a")
	assert.Equal(t, val, "ololo")
	assert.Nil(t, err)
}

func TestProcessCommand(t *testing.T) {
	st := CreateStorage()

	res, err := ProcessCommand(st, "SET a ololo")
	assert.Equal(t, "OK", res)
	assert.Nil(t, err)

	res, err = ProcessCommand(st, "GET a")
	assert.Equal(t, "ololo", res)
	assert.Nil(t, err)

	res, err = ProcessCommand(st, "DEL a")
	assert.Equal(t, "OK", res)
	assert.Nil(t, err)

	res, err = ProcessCommand(st, "GET a")

	assert.Equal(t, "", res)
	assert.Equal(t, errors.New("not found"), err)
}
