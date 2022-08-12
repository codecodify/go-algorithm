package chapter03

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func frequencySort(s string) string {
	if len(s) == 0 {
		return ""
	}

	var dict = make(map[string]int)
	for _, c := range s {
		dict[string(c)]++
	}

	var pairs = make(PairList, 0, len(dict))
	for k, v := range dict {
		pairs = append(pairs, Pair{k, v})
	}
	sort.Sort(pairs)
	var res string
	for _, p := range pairs {
		res += strings.Repeat(p.Key, p.Value)
	}
	return res
}
