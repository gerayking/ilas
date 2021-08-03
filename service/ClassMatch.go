package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/utils"
	"strconv"
)

const inf = 999999999

var g = &model.Graph{}

func Setg()  {
	g = global.Gragh
}

func CreateGraph(stu []model.Student, teacher []model.TeacherSchedule) {
	edges, head := make([]model.Edge, 0), make([]int, 0)
	LenOfHead, LenOfEdge := 0, 0
	for _, teacherSchedule := range teacher {
		key := strconv.Itoa(int(teacherSchedule.TeacherId))
		for _, index := range teacherSchedule.Schedule {
			k := key + index
			global.TeToIndex[k] = LenOfHead
			global.IndexToTe[LenOfHead] = k
			head = append(head, -1)
			head = append(head, -1)

			LenOfHead++
		}
	}
	nodeNumberOfTeacher := LenOfHead
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		for index, _ := range s.Plans {
			k := key + strconv.Itoa(index)
			global.StuToIndex[k] = LenOfHead
			global.IndexToStu[LenOfHead] = k
			head = append(head, -1)
			head = append(head,-1)
			LenOfHead++
		}
	}
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		for index, p := range s.Plans {
			keyOfStu := key + strconv.Itoa(index)
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
						edges = append(edges, model.Edge{W: 0, From: to, To: from, Next: head[to]})
						head[to] = LenOfEdge
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
	global.Gragh = &model.Graph{Head: head, Edges: edges, NodeNumber: LenOfHead, EdgeNumber: LenOfEdge, NodeNumberOfTeacher: nodeNumberOfTeacher}
}

// dfs分割
func dfs_to(v int, sonGraph *[]int, vis []bool) {
	vis[v]=true
	*sonGraph = append(*sonGraph, v)
	for i := g.Head[v]; i != -1; i = g.Edges[i].Next {
		t := g.Edges[i].To
		if !vis[t] {
			dfs_to(t, sonGraph, vis)
		}
	}
}

// 进行dfs将图进行分割
func Tarjan(n int) [][]int {
	vis := make([]bool,n)
	graphdivident := [][]int{}
	sgraph := make([]int, 0)
	for i := 0; i < n; i++ {
		if vis[i] == false {
			dfs_to(i, &sgraph, vis)
			graphdivident = append(graphdivident, sgraph)
			sgraph = make([]int, 0)
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
		for i :=head[u] ; i!=-1;i=g.Edges[i].Next {
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
	deep :=  make([]int,g.NodeNumber)
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

func addedge(u int, v int,w int) {
	g.Edges = append(g.Edges,model.Edge{W: w,From: u,To: v,Next: g.Head[u]})
	g.EdgeNumber++
}

func Match(multiGraph [][]int, n int) {
	superOriginNode := n
	superConvergeNode := superOriginNode + 1
	OriginNodeList := []int{}
	for i := 0; i < len(multiGraph); i++ {
		// 如果节点是学生+时间，连接上源点
		if len(multiGraph[i])<2 {
			continue
		}
		for j:=0;j<len(multiGraph[i]);j++{
			if multiGraph[i][j]< g.NodeNumberOfTeacher {
				 u := multiGraph[i][j]
				 addedge(u,superOriginNode,1)
				 addedge(superOriginNode,u,0)
			}else{
				u := multiGraph[i][j]
				addedge(u,superConvergeNode,1)
				addedge(superConvergeNode,u,0)
			}
		}
		OriginNodeList = append(OriginNodeList, superOriginNode)
		g.NodeNumber+=2
		superOriginNode += 2
		superConvergeNode += 2
		// 如果节点是学生+第几节课，连接上汇点
	}
	for _,item := range OriginNodeList{
		dinic(item, item+1)
	}
}

type Pair struct {
	First  int
	Second int
}

func OutputMatchInfo() []Pair {
	MatchInfo := make([]Pair,0)
	for u := 0; u < g.NodeNumberOfTeacher; u++ {
		for v := g.Head[u]; v != -1; v = g.Edges[v].Next {
			// 如果边的源点是学生计划节点则不匹配
			// 如果源点是老师放课节点且满流
			if v >= g.NodeNumberOfTeacher{
				continue
			}
			if g.Edges[v].W == 0 {
				MatchInfo = append(MatchInfo, Pair{u,v})
			}
		}
	}
	return MatchInfo
}
