package gio_test

import (
	"os"
	"testing"

	"github.com/MorrisMorrison/gutils/gio"
	"github.com/stretchr/testify/assert"
)

func TestReadLinesFromPathShouldReturnLines(t *testing.T) {
	want := []string{"8809887301", "5392059862"}
	lines, error := gio.ReadLinesFromPath("test_input.txt", true)
	if (error != nil){
		t.Fail();
	}

	assert.Equal(t, want, lines, "ReadLinesFromPath() output does not match the expected lines")
}



func TestReadLinesFromFileShouldReturnLines(t *testing.T) {
	file, err := os.Open("test_input.txt",)
	if err != nil {
		t.Fail()
	}
	defer file.Close()
	want := []string{"8809887301", "5392059862"}
	lines, error := gio.ReadLinesFromFile(file, true)
	if (error != nil){
		t.Fail();
	}

	assert.Equal(t, want, lines, "ReadLinesFromFile() output does not match the expected lines")
}
