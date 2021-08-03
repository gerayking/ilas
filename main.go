package main

import (
	"awesomeProject/global"
	"awesomeProject/service"
	"fmt"
)

func main() {
	sp,tp := CreateData()
	service.CreateGraph(sp,tp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph,global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	fmt.Println(ans)
}