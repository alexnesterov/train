package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	got := MaxInt(a, b)

	assert.Equal(t, b, got)
}

func TestMain(m *testing.M) {
	main()
}
