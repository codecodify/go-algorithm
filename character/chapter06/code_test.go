package chapter06

import (
	"fmt"
	"testing"
)

func TestBalancedStringSplit(t *testing.T) {
	fmt.Println(balancedStringSplit("RLLLRRRL"))
	fmt.Println(balancedStringSplit("LLLRRR"))
}
