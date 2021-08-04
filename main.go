package main

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/service"
	"fmt"
	"time"
)

func FirstMatch(sp []model.Student, tp []model.TeacherSchedule) []service.Pair {
	service.CreateGraph(sp, tp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	//service.MatchPlan2()
	ans := service.OutputMatchInfo()
	return ans

}
func SecondMatch(sp []model.Student, tp []model.TeacherSchedule) []service.Pair {
	service.CreateGraph(sp, tp)
	service.RebuildGraph(sp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	return service.OutputMatchInfo()
}

func testifyData(ans []service.Pair) {
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
func main() {
	start := time.Now()
	sp, tp := CreateData()
	ans := FirstMatch(sp, tp)
	testifyData(ans)
	ans = SecondMatch(sp, tp)
	testifyData(ans)
	end := time.Since(start)
	fmt.Println(end)
}
