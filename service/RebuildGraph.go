package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
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
							global.VirtualNode[keyOfTeacher] = append(global.VirtualNode[keyOfTeacher], global.StuToIndex[keyOfStu])
						} else {
							global.VirtualNode[keyOfTeacher] = make([]int, 0)
							global.VirtualNode[keyOfTeacher] = append(global.VirtualNode[keyOfTeacher], global.StuToIndex[keyOfStu])
						}

					}
				}
			}
		}
	}
	VitualNodeList := make([]TmpNode, 0)
	for k, v := range global.VirtualNode {
		VitualNodeList = append(VitualNodeList, TmpNode{TeacherId: k, LessonId: v})
	}
	sort.SliceStable(VitualNodeList, func(i, j int) bool {
		return len(VitualNodeList[i].LessonId) > len(VitualNodeList[j].LessonId)
	})
	for i := 0; i < 1000 && i < len(VitualNodeList); i++ {
		teacherId := VitualNodeList[i].TeacherId
		global.TeToIndex[teacherId] = global.Gragh.InitNumberOfNOde
		global.Gragh.Head = append(global.Gragh.Head, []int{-1, -1}...)
		u := global.Gragh.InitNumberOfNOde
		global.Gragh.NodeNumber++
		for _, item := range VitualNodeList[i].LessonId {
			addedge(u, item, 1)
			addedge(item, u, 0)
		}
	}
}
