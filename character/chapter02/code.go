package chapter02

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reorganize(s string) string {

	var dic = make(map[string]int)
	for _, v := range s {
		s := string(v)
		dic[s]++
	}
	// 由于go没有map的排序，所以需要自己实现一个排序，内置sort包是以切片为基础的排序
	var pairList = make(PairList, 0)
	for k, v := range dic {
		p := Pair{k, v}
		pairList = append(pairList, p)
	}
	// 排序
	sort.Sort(pairList)
	if len(s) == 0 || pairList[len(pairList)-1].Value > (len(s)+1)/2 {
		return ""
	}
	//fmt.Printf("%v\n", pairList)

	// 当前要插入的字符
	var pair Pair

	// 开始执行奇偶插空
	var res = make([]string, len(s))
	for i := 0; i < len(s); i += 2 {
		if pair.Value == 0 {
			// 如果是Java或Python等有内置数组操作方法,可以使用pop()方法
			pair = pairList[len(pairList)-1]
			pairList = pairList[:len(pairList)-1]
		}
		res[i] = pair.Key
		pair.Value--
	}
	for i := 1; i < len(s); i += 2 {
		if pair.Value == 0 {
			pair = pairList[len(pairList)-1]
			pairList = pairList[:len(pairList)-1]
		}
		res[i] = pair.Key
		pair.Value--
	}
	return strings.Join(res, "")
}
