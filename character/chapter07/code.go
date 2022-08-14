package chapter07

// 判断是否回文
func isPal(s string) bool {
	return s == reverse(s)
}

// 字符串倒序
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 递归进行分割
func clipStr(start int, s []rune, l [][]rune, res [][][]rune) [][][]rune {
	// 如果已经到达字符串的末尾，则将结果添加到结果集中
	if start > len(string(s))-1 {
		res = append(res, l)
		return res
	}

	for i := start + 1; i < len(s)+1; i++ {
		// 如果是回文，递归继续分割
		if isPal(string(s[start:i])) {
			res = clipStr(i, s, append(l, s[start:i]), res)
		}
	}
	return res
}

func partition(str string) [][]string {
	var res [][][]rune
	res = clipStr(0, []rune(str), [][]rune{}, [][][]rune{})
	var ret [][]string
	for _, v := range res {
		var s []string
		for _, vv := range v {
			s = append(s, string(vv))
		}
		ret = append(ret, s)
	}
	return ret
}
