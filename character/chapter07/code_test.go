package chapter07

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(reverse("abc"))
}

func TestPartition(t *testing.T) {
	fmt.Println(partition("abb"))
}
