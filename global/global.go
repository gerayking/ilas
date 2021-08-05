package global

import "awesomeProject/model"

var Gragh *model.Graph
var TeToIndex, StuToIndex, IndexToTe, IndexToStu = make(map[string]int), make(map[string]int), make(map[int]string), make(map[int]string)
var VirtualNode = make(map[string][]int)
var StuToPlan = make(map[string][]model.Plan)
var StuToTe = make(map[string][]uint)
var InitNumberOfNOde int
var InDegree []int
var InFirstMatch []bool
var ReceiveChanel = make(chan model.TeacherPlan)
var RequestChanel = make(chan model.TeacherPlan)
var RemainStudentPlan []model.Student
var RemainTeacherPlan []model.TeacherPlan
var IsMatchSchedule map[string]int
