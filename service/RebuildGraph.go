package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/utils"
	"sort"
	"strconv"
	"strings"
)

// 添加1000条边后重新建图
func RebuildGraph(stu []model.Student) {
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
						_, ok := global.VirtualNode[keyOfTeacher]
						if ok {
							if utils.Contains(global.VirtualNode[keyOfTeacher], global.StuToIndex[keyOfStu]) == false {
								global.VirtualNode[keyOfTeacher] = append(global.VirtualNode[keyOfTeacher], global.StuToIndex[keyOfStu])
							}
						} else {
							global.VirtualNode[keyOfTeacher] = make([]int, 0)
							global.VirtualNode[keyOfTeacher] = append(global.VirtualNode[keyOfTeacher], global.StuToIndex[keyOfStu])
						}
					}
				}
			}
		}
	}
	VirtualNodeList := make([]model.TmpNode, 0)
	for k, v := range global.VirtualNode {
		VirtualNodeList = append(VirtualNodeList, model.TmpNode{TeacherId: k, LessonId: v})
	}
	sort.SliceStable(VirtualNodeList, func(i, j int) bool {
		return len(VirtualNodeList[i].LessonId) > len(VirtualNodeList[j].LessonId)
	})
	for i := 0; i < len(VirtualNodeList) && i < 1000; i++ {
		teacherId := VirtualNodeList[i].TeacherId
		global.TeToIndex[teacherId] = global.Gragh.NodeNumber
		global.IndexToTe[global.Gragh.NodeNumber] = teacherId
		global.Gragh.Head = append(global.Gragh.Head, []int{-1, -1}...)
		u := global.Gragh.NodeNumber
		for _, item := range VirtualNodeList[i].LessonId {
			utils.Addedge(item, u, 1)
			utils.Addedge(u, item, 0)
		}
		global.Gragh.NodeNumber++
	}
}

// 获取匹配后的残留图
func GetRemainGraph(ans []model.Pair) ([]model.Student, []model.TeacherPlan) {
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
			ids := strings.Split(Sid, "_")[0]
			id, _ := strconv.ParseInt(ids, 0, 0)
			student := model.Student{StuId: uint(id), Plans: global.StuToPlan[ids], Teachers: global.StuToTe[ids]}
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
