package main

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/service"
	"strconv"
	"strings"
)

func FirstMatch(sp []model.Student, tp []model.TeacherPlan) []service.Pair {
	service.CreateGraph(sp, tp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	return ans

}
func SecondMatch(sp []model.Student, tp []model.TeacherPlan) []service.Pair {
	service.CreateGraph(sp, tp)
	service.RebuildGraph(sp)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	return service.OutputMatchInfo()
}
func getRemainGraph(ans []service.Pair) ([]model.Student, []model.TeacherPlan) {
	for i := 0; i < len(ans); i++ {
		u := ans[i].First
		v := ans[i].Second
		s := global.IndexToStu[u]
		t := global.IndexToTe[v]
		delete(global.IndexToStu, u)
		delete(global.IndexToTe, v)
		delete(global.StuToIndex, s)
		delete(global.TeToIndex, t)
	}
	global.Gragh = &model.Graph{}
	studentPlan := make([]model.Student, 0)
	teacherPlan := make([]model.TeacherPlan, 0)
	for _, item := range global.StuToIndex {
		if global.InFirstMatch[item] == false {
			Sid := global.IndexToStu[item]
			id, _ := strconv.ParseInt(strings.Split(Sid, "_")[0], 0, 0)
			student := model.Student{StuId: uint(id), Plans: global.StuToPlan[global.IndexToStu[item]], Teachers: global.StuToTe[global.IndexToStu[item]]}
			studentPlan = append(studentPlan, student)
		}
	}
	for _, item := range global.TeToIndex {
		if global.InFirstMatch[item] == false {
			ts := strings.Split(global.IndexToTe[item], "_")
			teacherId, _ := strconv.ParseInt(ts[0], 0, 0)
			teacherPlan = append(teacherPlan, model.TeacherPlan{TeacherId: teacherId, Schedule: []string{ts[1]}})
		}
	}
	return studentPlan, teacherPlan
}
func main() {
	//start := time.Now()
	sp, tp := service.CreateData()
	FirstMatch(sp, tp)
	ans := SecondMatch(sp, tp)
	global.RemainStudentPlan, global.RemainTeacherPlan = getRemainGraph(ans)
	//end := time.Since(start)
	//fmt.Println(end)
	//testifyData(ans)
	//testifyData(ans)
	//end = time.Since(start)
	//fmt.Println(end)
}
