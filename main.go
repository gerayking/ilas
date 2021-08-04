package main

import (
	"awesomeProject/global"
	"awesomeProject/service"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sp,tp := CreateData()
	service.CreateGraph(sp,tp)
	service.Setg()
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph,global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()


	for i:=0;i< len(ans);i++{
		u := ans[i].First
		v := ans[i].Second
		s := global.IndexToStu[v]
		t := global.IndexToTe[u]
		//fmt.Println(strconv.Itoa(u) +"--------" + strconv.Itoa(v))
		fmt.Println(s +"--------" + t)


	}
	fmt.Println(len(ans))
	end := time.Since(start)
	fmt.Println(end)
}