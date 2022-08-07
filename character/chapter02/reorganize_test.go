package chapter02

import (
	"fmt"
	"testing"
)

func Test_reorganize(t *testing.T) {
	s := reorganize("aab")
	fmt.Println(s)

	s = reorganize("aaab")
	fmt.Println(s)

	s = reorganize("bbaacc")
	fmt.Println(s)
}
