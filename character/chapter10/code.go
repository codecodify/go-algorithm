package chapter10

func firstUniqueChar(s string) string {
	// 记录字符是否出现
	var hash = make(map[string]int)
	// 定义字符列表
	var words []string
	for i, v := range s {
		k := string(v)
		if _, ok := hash[k]; ok {
			words = remove(words, k)
		} else {
			words = append(words, k)
			hash[k] = i
		}
	}
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

// 移除元素
func remove(s []string, w string) []string {
	for i, v := range s {
		if v == w {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
