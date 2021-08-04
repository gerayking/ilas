package utils

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/service"
	"fmt"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func Addedge(u int, v int, w int) {
	global.Gragh.Edges = append(global.Gragh.Edges, model.Edge{W: w, From: u, To: v, Next: global.Gragh.Head[u]})
	global.Gragh.Head[u] = global.Gragh.EdgeNumber
	global.Gragh.EdgeNumber++
}
func TestifyAndOutpurData(ans []service.Pair) {
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
		flag := service.IsMatch(&ans[i])
		if flag {
			trueNumber++
		}

	}
	fmt.Printf("Match number : %d\n", len(ans))
	fmt.Printf("success Match Number : %d\n", trueNumber)
	fmt.Printf("error number : %d\n", confilictNumber)
}
