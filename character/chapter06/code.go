package chapter06

func balancedStringSplit(str string) (n int) {
	// 分别记录'R'和'L'的个数
	var r, l int
	for _, v := range str {
		if v == 'R' {
			r++
		} else if v == 'L' {
			l++
		}
		if r == l {
			n++
			r = 0
			l = 0
		}
	}
	return
}