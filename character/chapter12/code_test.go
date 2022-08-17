package chapter12

import (
	"fmt"
	"testing"
)

func TestCommonChars(t *testing.T) {
	fmt.Println(commonChars([]string{"ball", "gad", "apple"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
