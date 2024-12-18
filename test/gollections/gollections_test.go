package gollections_test

import (
	"testing"

	"github.com/MorrisMorrison/gutils/gollections"
	"github.com/go-playground/assert/v2"
)

func TestAnyShouldReturnTrue(t *testing.T) {
	names := []string{"Morris", "Morrison"}
	want := true
	actual := gollections.Any(names, func(x string) bool {return x == "Morris" })

	assert.Equal(t, want, actual)
}

func TestAnyShouldReturnFalse(t *testing.T) {
	names := []string{"Morrisun", "Morrison"}
	want := true
	actual := gollections.Any(names, func(x string) bool {return x == "Morris" })
	assert.NotEqual(t, want, actual)
}

func TestAllShouldReturnTrue(t *testing.T){
	names := []string{"Morris", "Morris"}
	want := true
	actual := gollections.All(names, func(x string) bool {return x == "Morris" })

	assert.Equal(t, want, actual)
}

func TestAllShouldReturnFalse(t *testing.T) {
	names := []string{"Morris", "Morrison"}
	want := true
	actual := gollections.All(names, func(x string) bool {return x == "Morris" })
	assert.NotEqual(t, want, actual)
}


func TestEqualsShouldReturnTrue(t *testing.T){
	names := []string{"Morris", "Morrisa"}
	compareTo:= []string{"Morris", "Morrisa"}
	want := true
	actual := gollections.Equals(names, compareTo, func(x string, y string) bool {return x == y })

	assert.Equal(t, want, actual)
}

func TestEqualsShouldReturnFalse(t *testing.T) {
	names := []string{"Morris", "Morrison"}
	compareTo:= []string{"Morris", "Morrisa"}
	want := false
	actual := gollections.Equals(names,compareTo, func(x string, y string) bool {return x == y })
	assert.Equal(t, want, actual)
}