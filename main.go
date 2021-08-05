package main

import (
	"awesomeProject/Random"
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"time"
)

func PreRun(tp []model.TeacherPlan) {
	_, _, global.IsMatchSchedule = Random.GetTimeCount(tp)
	go Random.EmitData()
	go Random.OnQuantityOrTimeMatch()
}
func FirstMatch(sp []model.Student, tp []model.TeacherPlan) []model.Pair {
	service.CreateGraph(sp, tp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	utils.TestifyAndOutpurData(ans)
	return ans

}
func SecondMatch(sp []model.Student, tp []model.TeacherPlan) []model.Pair {
	service.CreateGraph(sp, tp)
	service.RebuildGraph(sp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	utils.TestifyAndOutpurData(ans)
	return service.OutputMatchInfo()
}
func main() {
	start := time.Now()
	sp, tp := service.CreateData()
	//PreRun(tp) // 启动候补进程

	FirstMatch(sp, tp) // 第一次匹配求出1000条可行边
	end := time.Since(start)
	fmt.Println(end)
	start = time.Now()
	ans := SecondMatch(sp, tp) // 第二次匹配求出最终值
	utils.TestifyAndOutpurData(ans)
	end = time.Since(start)
	fmt.Println(end)
	//global.RemainStudentPlan, global.RemainTeacherPlan = service.GetRemainGraph(ans)
	//Random.StimulateTeacher(sp) // 刺激放课
	//end := time.Since(start)
	//fmt.Println(end)
	//testifyData(ans)
	//testifyData(ans)
	//end = time.Since(start)
	//fmt.Println(end)
}
