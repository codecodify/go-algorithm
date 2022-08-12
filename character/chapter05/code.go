package chapter05

import (
	"fmt"
	"strings"
)

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}

	var dic = make(map[string][]int)
	for i, v := range s {
		k := string(v)
		dic[k] = append(dic[k], i)
	}

	fmt.Println(dic)

	// 由于第二个字符串不是通过range循环得到的，所以需要转换成[]rune，否则中文或其它特殊字符就匹配不上
	runes := []rune(t)

	for _, indexes := range dic {
		// 对应第二个字符串指定位置的字符
		w := runes[indexes[0]]

		if len(indexes) != strings.Count(t, string(w)) {
			return false
		}
		if len(indexes) > 1 {
			for _, i := range indexes {
				if w != runes[i] {
					return false
				}
			}
		}
	}

	return true
}
