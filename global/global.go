package global

import "awesomeProject/model"

var Gragh *model.Graph //图
// 节点编号到老师学生信息的映射
var TeToIndex, StuToIndex, IndexToTe, IndexToStu = make(map[string]int), make(map[string]int), make(map[int]string), make(map[int]string)
var VirtualNode = make(map[string][]int)      // 虚拟节点
var StuToPlan = make(map[string][]model.Plan) // 学生计划映射
var StuToTe = make(map[string][]uint)         // 学生到老师列表映射
var InitNumberOfNOde int
var InDegree []int
var InFirstMatch []bool                          // 第一次匹配成功的节点编号
var ReceiveChanel = make(chan model.TeacherPlan) // 用于候补缓解的模拟
var RequestChanel = make(chan model.TeacherPlan) // 用于候补缓解的模拟
var RemainStudentPlan []model.Student            // 残留学生计划
var RemainTeacherPlan []model.TeacherPlan        // 残留老师计划
var IsMatchSchedule map[string]int               // 是否匹配
var MatchingFlag chan int
