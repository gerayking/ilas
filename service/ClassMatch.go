package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/utils"
	"container/list"
	"strconv"
)

const inf = 999999999

var g = global.Gragh

func CreateGraph(stu []model.Student, teacher []model.TeacherSchedule) {
	edges := make([]model.Edge, 0)
	LenOfHead, LenOfEdge := 0, 0
	for _, teacherSchedule := range teacher {
		key := strconv.Itoa(int(teacherSchedule.TeacherId))
		for _, index := range teacherSchedule.Schedule {
			key += index
			global.TeToIndex[key] = LenOfHead
			global.IndexToTe[LenOfHead] = key
			LenOfHead++
		}
	}
	nodeNumberOfTeacher := LenOfHead
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		for index, _ := range s.Plans {
			key += strconv.Itoa(index)
			global.StuToIndex[key] = LenOfHead
			global.IndexToStu[LenOfHead] = key
			LenOfHead++
		}
	}
	head := make([]int, LenOfHead)
	for _, s := range stu {
		keyOfStu := strconv.Itoa(int(s.StuId))
		for index, p := range s.Plans {
			keyOfStu += strconv.Itoa(index)
			for _, c := range p.Class {
				for _, t := range s.Teachers {
					keyOfTeacher := strconv.Itoa(int(t))
					keyOfTeacher += c
					from := global.StuToIndex[keyOfStu]
					to, ok := global.TeToIndex[keyOfTeacher]
					if ok {
						edges = append(edges, model.Edge{W: 1, From: from, To: to, Next: head[from]})
						head[from] = LenOfEdge
						LenOfEdge++
					} else {
						_, ok := global.VirtualNode[keyOfTeacher]
						if !ok {
							global.VirtualNode[keyOfTeacher] = len(global.InDegree)
							global.InDegree = append(global.InDegree, 1)
						}
						global.InDegree[global.VirtualNode[keyOfTeacher]]++
					}
				}
			}
		}
	}
	g = &model.Graph{Head: head, Edges: edges, NodeNumber: LenOfHead, EdgeNumber: LenOfEdge, NodeNumberOfTeacher: nodeNumberOfTeacher}
}

// dfs分割
func dfs_to(v int, sonGraph list.List, vis []bool) {
	sonGraph.PushBack(v)
	for i := g.Head[v]; i != -1; i = g.Edges[i].Next {
		t := g.Edges[i].To
		if !vis[t] {
			dfs_to(t, sonGraph, vis)
		}
		sonGraph.Init()
	}
}

// 进行dfs将图进行分割
func Tarjan(n int) list.List {
	vis := []bool{}
	graphdivident := list.List{}
	sgraph := list.List{}
	for i := 0; i < n; i++ {
		if vis[i] == false {
			dfs_to(i, sgraph, vis)
		}
		graphdivident.PushBack(sgraph)
		var next *list.Element
		for e := sgraph.Front(); e != nil; e = next {
			next = e.Next()
			sgraph.Remove(e)
		}
	}
	return graphdivident
}

// dinic算法分层
func bfs(edges []model.Edge, deep []int, head []int, s int, t int) int {
	queue := utils.New()
	queue.Push(0)
	for queue.Len() != 0 {
		queue.Pop()
	}
	deep[s] = 1
	queue.Push(s)
	for queue.Len() != 0 {
		u := queue.Peek().(int)
		queue.Pop()
		i := 0
		for i != -1 {
			i = head[u]
			if edges[i].W > 0 && deep[edges[i].To] == 0 {
				deep[edges[i].To] = deep[edges[i].From] + 1
				queue.Push(edges[i].To)
			}
			i = edges[i].Next
		}
	}
	if deep[t] == 0 {
		return 0
	}
	return 1
}

// 寻找增广路
func dfs(edges []model.Edge, head []int, deep []int, u int, t int, dist int) int {
	if u == t {
		return dist
	}
	i := 0
	for i != -1 {
		i = head[u]
		if deep[edges[i].To] == deep[edges[i].From]+1 && edges[i].W != 0 {
			di := dfs(edges, head, deep, edges[i].To, t, dist)
			if di > 0 {
				edges[i].W -= di
				edges[i^1].W += di
				return di
			}
		}
	}
	return 0
}

// dinic 计算最大流
func dinic(u int, v int) int {
	ans := 0
	var deep []int
	// 对残流图不断进行分层
	for bfs(g.Edges, deep, g.Head, u, v) != 0 {
		// 分层后寻找增广路
		minflow := dfs(g.Edges, g.Head, deep, u, v, inf)
		if minflow != 0 {
			ans += minflow
		}
	}
	return ans
}

func addedge(u int, v int) {
	g.Edges[g.EdgeNumber].W = 1
	g.Edges[g.EdgeNumber].From = u
	g.Edges[g.EdgeNumber].To = v
	g.Edges[g.EdgeNumber].Next = g.Head[u]
	g.EdgeNumber++
}

func Match(multiGraph list.List, n int) {
	superOriginNode := n
	superConvergeNode := superOriginNode + 1
	OriginNodeList := list.List{}
	for item := multiGraph.Front(); item != nil; item = item.Next() {
		// 如果节点是老师+时间，连接上源点
		if true {
			addedge(superOriginNode, item.Value.(int))
		} else {
			addedge(item.Value.(int), superConvergeNode)
		}
		OriginNodeList.PushFront(superOriginNode)
		superOriginNode += 2
		superConvergeNode += 2
		// 如果节点是学生+第几节课，连接上汇点
	}
	for item := OriginNodeList.Front(); item != nil; item = item.Next() {
		superOriginNode := item.Value.(int)
		superConvergeNode := superOriginNode + 1
		dinic(superOriginNode, superConvergeNode)
	}
}

type Pair struct {
	first  int
	second int
}

func OutputMatchInfo() list.List {
	MatchInfo := list.List{}
	for u := 0; u < g.NodeNumberOfTeacher; u++ {
		for v := g.Head[u]; v != -1; v = g.Edges[v].Next {
			// 如果边的源点是学生计划节点则不匹配
			// 如果源点是老师放课节点且满流
			if v >= g.NodeNumberOfTeacher{
				continue
			}
			if g.Edges[v].W == 0 {
				MatchInfo.PushFront(Pair{u, v})
			}
		}
	}
	return MatchInfo
}
