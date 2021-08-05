package utils

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"fmt"
	"strconv"
	"strings"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// 临接表进行图加边
func Addedge(u int, v int, w int) {
	global.Gragh.Edges = append(global.Gragh.Edges, model.Edge{W: w, From: u, To: v, Next: global.Gragh.Head[u]})
	global.Gragh.Head[u] = global.Gragh.EdgeNumber
	global.Gragh.EdgeNumber++
}

// 测试匹配是否正确
func IsMatch(p *model.Pair) bool {
	stu := strings.Split(global.IndexToStu[p.First], "_")
	te := strings.Split(global.IndexToTe[p.Second], "_")
	i, _ := strconv.ParseInt(stu[1], 0, 0)
	flag := false
	for _, plan := range global.StuToPlan[stu[0]][i].Class {
		if strings.EqualFold(plan, te[1]) {
			flag = true
		}
	}
	if flag {
		for _, teacher := range global.StuToTe[stu[0]] {
			if strings.EqualFold(te[0], strconv.Itoa(int(teacher))) {
				return true
			}
		}
	}
	return false
}

// 测试匹配结果的正确率以及正确性
func TestifyAndOutpurData(ans []model.Pair) {
	trueNumber := 0
	confilictNumber := 0
	mps := make(map[string]int)
	mpt := make(map[string]int)
	for i := 0; i < len(ans); i++ {
		u := ans[i].First
		v := ans[i].Second
		s := global.IndexToStu[u]
		t := global.IndexToTe[v]
		if _, ok := mps[s]; ok && mps[s] == 1 {
			confilictNumber++
		} else {
			mps[s] = 1
		}
		if _, ok := mps[t]; ok && mpt[t] == 1 {
			confilictNumber++
		} else {
			mpt[t] = 1
		}
		flag := IsMatch(&ans[i])
		if flag {
			trueNumber++
		}

	}
	fmt.Printf("Match number : %d\n", len(ans))
	fmt.Printf("success Match Number : %d\n", trueNumber)
	fmt.Printf("error number : %d\n", confilictNumber)
}
