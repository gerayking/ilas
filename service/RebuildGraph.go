package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/utils"
	"sort"
	"strconv"
)

type TmpNode struct {
	TeacherId string
	LessonId  []int
}

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
	VirtualNodeList := make([]TmpNode, 0)
	for k, v := range global.VirtualNode {
		VirtualNodeList = append(VirtualNodeList, TmpNode{TeacherId: k, LessonId: v})
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
			addedge(item, u, 1)
			addedge(u, item, 0)
		}
		global.Gragh.NodeNumber++
	}
}
