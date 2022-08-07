package chapter01

import (
	"math"
	"strings"
)

func reformat(s string) (res string) {
	numStr := "0123456789"

	var numList, charList []string

	for _, v := range s {
		str := string(v)
		if strings.Contains(numStr, string(v)) {
			numList = append(numList, str)
		} else {
			charList = append(charList, str)
		}
	}
	if math.Abs(float64(len(numList)-len(charList))) > 1 {
		return
	}
	if len(numList) > len(charList) {
		for i := 0; i < len(charList); i++ {
			res += numList[i]
			res += charList[i]
		}
		res += numList[len(charList)]
	} else if len(numList) == len(charList) {
		for i := 0; i < len(charList); i++ {
			res += numList[i]
			res += charList[i]
		}
	} else {
		for i := 0; i < len(numList); i++ {
			res += charList[i]
			res += numList[i]
		}
		res += charList[len(numList)]
	}
	return res
}
