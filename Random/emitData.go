package Random

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func GetTimeCount(tp []model.TeacherPlan) ([]string, map[string]int, map[string]int) {
	TimeToIndex := make(map[string]int)
	TimeCount := make(map[string]int)
	count := 0
	for _, item := range tp {
		for _, s := range item.Schedule {
			if _, ok := TimeToIndex[s]; !ok {
				TimeToIndex[s] = count
				count++
			}
			if _, ok := TimeCount[s[1:]]; !ok {
				TimeCount[s[1:]] = 1
			} else {
				TimeCount[s[1:]]++
			}
		}
	}
	IndexToTime := make([]string, len(TimeToIndex))
	for key, value := range TimeToIndex {
		IndexToTime[value] = key
	}
	sort.Slice(IndexToTime, func(i, j int) bool {
		return IndexToTime[i] < IndexToTime[j]
	})
	return IndexToTime, TimeToIndex, TimeCount
}

func IsEmit(schedule string) bool {
	k := global.IsMatchSchedule
	rf := rand.Float32()
	if k[schedule] < 200 {
		return rf < 0.3
	} else if k[schedule] < 4000 {
		return rf < 0.4
	} else if k[schedule] < 7000 {
		return rf < 0.5
	} else if k[schedule] < 20000 {
		return rf < 0.6
	} else {
		return rf < 0.7
	}
}
func StimulateTeacher(stu []model.Student) {
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		for index, p := range s.Plans {
			keyOfStu := key + "_" + strconv.Itoa(index)
			for _, c := range p.Class {
				for _, t := range s.Teachers {
					keyOfTeacher := strconv.Itoa(int(t))
					keyOfTeacher += "_" + c
					_, ok := global.TeToIndex[keyOfTeacher]
					if global.InFirstMatch[global.TeToIndex[keyOfTeacher]] == true || global.InFirstMatch[global.StuToIndex[keyOfStu]] == true {
						continue
					}
					if !ok {
						global.RequestChanel <- model.TeacherPlan{TeacherId: int64(t), Schedule: []string{c}}
					}
				}
			}
		}
	}
}
func EmitData() {
	for true {
		me := <-global.RequestChanel
		for _, item := range me.Schedule {
			if IsEmit(item) {
				global.ReceiveChanel <- me
			}
		}
	}
}
func ReMatch(teacherPlanList []model.TeacherPlan, c chan int) {
	global.RemainTeacherPlan = append(global.RemainTeacherPlan, teacherPlanList...)
	service.CreateGraph(global.RemainStudentPlan, global.RemainTeacherPlan)
	sonGraph := service.Tarjan(global.Gragh.NodeNumber)
	service.Match(sonGraph, global.Gragh.NodeNumber)
	ans := service.OutputMatchInfo()
	utils.TestifyAndOutpurData(ans)
	global.RemainStudentPlan, global.RemainTeacherPlan = service.GetRemainGraph(ans)
	c <- 1
}

/*i8
定时， 2定量
*/
func OnQuantityOrTimeMatch() {
	planList := make([]model.TeacherPlan, 0)
	timer1 := time.NewTimer(time.Second * 10)
	go func(plan *[]model.TeacherPlan) {
		<-timer1.C
		fmt.Println("Time is able to Run")
		ci := make(chan int)
		ReMatch(planList, ci)
		<-ci
		planList = make([]model.TeacherPlan, 0)
	}(&planList)
	for true {
		me := <-global.ReceiveChanel
		planList = append(planList, me)
		if len(planList) > 1000 {
			fmt.Println("Quantity is able to Run")
			ci := make(chan int)
			ReMatch(planList, ci)
			<-ci
			planList = make([]model.TeacherPlan, 0)
			timer1.Reset(time.Second * 30)
		}
	}
}
