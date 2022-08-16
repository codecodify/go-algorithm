package chapter11

func maxL(start int, str string) (max int) {
	var hash = make(map[string]int)
	// 这里需要转换[]rune，否则直接用str切片(str[start:])会乱码
	runes := []rune(str)
	for i, s := range runes[start:] {
		if _, ok := hash[string(s)]; ok {
			return
		}
		hash[string(s)] = i
		max++
	}
	return
}

func lengthOfLongestSubstring(str string) (max int) {
	for i, _ := range str {
		length := maxL(i, str)
		if length > max {
			max = length
		}
	}
	return
}
