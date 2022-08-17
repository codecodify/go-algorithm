package chapter12

func commonChars(str []string) (res []string) {
	if len(str) == 0 {
		return
	}
	var hash = make(map[string]int)
	for i, v := range str[0] {
		s := string(v)
		hash[s] = i
	}
	for _, v := range str[1:] {
		var swap = make(map[string]int)
		for _, v := range v {
			s := string(v)
			if i, ok := hash[s]; ok {
				swap[s] = i
			}
		}
		hash = swap
	}

	for k := range hash {
		res = append(res, k)
	}
	return
}
