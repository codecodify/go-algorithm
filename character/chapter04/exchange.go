package chapter04

func exchange(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	n, m := -1, -1
	for i := range a {
		if a[i] != b[i] {
			if n == -1 {
				n = i
			}else if m == -1 {
				m = i
			}else {
				return false
			}
		}
	}

	if a == b && len(set(a)) < len(a) {
		return true
	}

	if n != -1 && m != -1 && a[n] == b[m] && a[m] == b[n] {
		return true
	}

	return false
}

// 字符串去重
func set(s string) string {
	var hashTable = make(map[string]string)
	var res string
	for _, v := range s {
		k := string(v)
		if _, ok := hashTable[k]; !ok {
			hashTable[k] = k
			res += string(v)
		}
	}
	return res
}
