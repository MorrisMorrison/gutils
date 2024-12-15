package gio_test

import (
	"testing"

	"github.com/MorrisMorrison/gutils/gio"
	"github.com/stretchr/testify/assert"
)

func TestReadFileShouldReturnLines(t *testing.T) {
	want := []string{"8809887301", "5392059862"}
	lines := gio.ReadFile("test_input.txt", true)

	assert.Equal(t, want, lines, "ReadFile() output does not match the expected lines")
}
