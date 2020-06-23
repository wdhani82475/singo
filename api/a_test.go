package api

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T)  {
	aa(3)
}

func aa(b int) int {
	fmt.Println(b)
	return b*b
}
