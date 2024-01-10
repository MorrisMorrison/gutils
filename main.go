package main

import (
	"fmt"

	"github.com/MorrisMorrison/gutils/gollections"
)

func main() {
	valuesToCheck := []int{0, 1, 2, 3, 4, 5}
	containsValue := gollections.Contains(valuesToCheck, 1)
	fmt.Println(containsValue)
}
