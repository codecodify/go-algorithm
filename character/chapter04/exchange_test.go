package chapter04

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {
	fmt.Println(exchange("acb", "cab"))

	fmt.Println(exchange("aaa", "bbb"))
}

func TestSet(t *testing.T) {
	fmt.Println(set("aabbccd"))
	fmt.Println(set("aaaa"))
}
