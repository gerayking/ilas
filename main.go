package main

import (
	"awesomeProject/global"
	"awesomeProject/service"
	"fmt"
)

func main() {
	sp,tp := CreateData()
	service.CreateGraph(sp,tp)
	service.Setg()
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph,global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	fmt.Println(ans)
	fmt.Println(len(ans))
	//for i:=0;i<len(ans);i++{
	//	u := ans[i].First
	//	v := ans[i].Second
	//	s := global.IndexToStu[u]
	//	t := global.IndexToTe
	//}
}