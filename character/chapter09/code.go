package chapter09

func maxPower(str string) (max int) {

	// 当前循环的字符
	var c string

	// 当前循环的字符个数
	var count int

	for _, v := range str {
		s := string(v)
		if c == "" {
			c = s
			count = 1
		} else {
			if c == s {
				count++
			} else {
				c = s
				count = 1
			}
		}
		if count > max {
			max = count
		}
	}

	return
}
