package main

import (
	"awesomeProject/global"
	"awesomeProject/service"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sp, tp := CreateData()
	service.CreateGraph(sp, tp)
	service.Setg()
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	fmt.Println(len(ans))
	service.RebuildGraph(sp)
	service.Setg()
	sonGraph = service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans = service.OutputMatchInfo()
	fmt.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		u := ans[i].First
		v := ans[i].Second
		s := global.IndexToStu[u]
		t := global.IndexToTe[v]
		//fmt.Println(strconv.Itoa(u) +"--------" + strconv.Itoa(v))
		fmt.Println(s + "--------" + t)

	}
	end := time.Since(start)
	fmt.Println(end)
}
