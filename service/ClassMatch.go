package service

import (
	"awesomeProject/global"
	"awesomeProject/model"
	"awesomeProject/utils"
	"math"
	"strconv"
	"strings"
	"sync"
)

const inf = 999999999

func CreateGraph(stu []model.Student, teacher []model.TeacherSchedule) {
	edges, head := make([]model.Edge, 0), make([]int, 0)
	LenOfHead, LenOfEdge := 0, 0
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		global.StuToPlan[key] = s.Plans
		global.StuToTe[key] = s.Teachers
		for index, _ := range s.Plans {
			k := key + "_" + strconv.Itoa(index)
			global.StuToIndex[k] = LenOfHead
			global.IndexToStu[LenOfHead] = k
			head = append(head, -1)
			head = append(head, -1)
			LenOfHead++
		}
	}
	nodeNumberOfStu := LenOfHead
	for _, teacherSchedule := range teacher {
		key := strconv.Itoa(int(teacherSchedule.TeacherId))
		for _, index := range teacherSchedule.Schedule {
			k := key + "_" + index
			global.TeToIndex[k] = LenOfHead
			global.IndexToTe[LenOfHead] = k
			head = append(head, -1)
			head = append(head, -1)
			LenOfHead++
		}
	}
	for _, s := range stu {
		key := strconv.Itoa(int(s.StuId))
		for index, p := range s.Plans {
			keyOfStu := key + "_" + strconv.Itoa(index)
			for _, c := range p.Class {
				for _, t := range s.Teachers {
					keyOfTeacher := strconv.Itoa(int(t))
					keyOfTeacher += "_" + c
					from := global.StuToIndex[keyOfStu]
					to, ok := global.TeToIndex[keyOfTeacher]
					if ok {
						edges = append(edges, model.Edge{W: 1, From: from, To: to, Next: head[from]})
						head[from] = LenOfEdge
						LenOfEdge++
						edges = append(edges, model.Edge{W: 0, From: to, To: from, Next: head[to]})
						head[to] = LenOfEdge
						LenOfEdge++
					}
				}
			}
		}
	}
	global.Gragh = &model.Graph{Head: head, Edges: edges, NodeNumber: LenOfHead, EdgeNumber: LenOfEdge, NodeNumberOfStu: nodeNumberOfStu}
}

// dfs分割
func dfs_to(v int, sonGraph *[]int, vis []bool) {
	vis[v] = true
	*sonGraph = append(*sonGraph, v)
	for i := global.Gragh.Head[v]; i != -1; i = global.Gragh.Edges[i].Next {
		t := global.Gragh.Edges[i].To
		if !vis[t] {
			dfs_to(t, sonGraph, vis)
		}
	}
}

// 进行dfs将图进行分割
func Tarjan(n int) [][]int {
	vis := make([]bool, n)
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
func CreateSuperNode(n int) {

}

// dinic算法分层
func bfs(edges []model.Edge, deep []int, head []int, s int, t int, qu *[]int) int {
	queue := utils.New()
	for queue.Len() != 0 {
		queue.Pop()
	}
	deep[s] = 1
	queue.Push(s)
	for queue.Len() != 0 {
		u := queue.Peek().(int)
		*qu = append(*qu, u)
		queue.Pop()
		for i := head[u]; i != -1; i = global.Gragh.Edges[i].Next {
			if edges[i].W > 0 && deep[edges[i].To] == 0 {
				deep[edges[i].To] = deep[edges[i].From] + 1
				queue.Push(edges[i].To)
			}
		}
	}
	if deep[t] == 0 {
		return 0
	}
	return 1
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 寻找增广路
func dfs(edges []model.Edge, head []int, deep []int, vis []bool, u int, t int, dist int) int {
	if u == t {
		return dist
	}
	for i := head[u]; i != -1; i = global.Gragh.Edges[i].Next {
		if deep[edges[i].To] == deep[edges[i].From]+1 && edges[i].W != 0 {
			di := dfs(edges, head, deep, vis, edges[i].To, t, min(dist, edges[i].W))
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
func dinic(u int, v int, group *sync.WaitGroup) int {

	ans := 0
	deep := make([]int, 2*global.Gragh.NodeNumber)
	vis := make([]bool, 2*global.Gragh.NodeNumber)
	// 对残流图不断进行分层
	qu := make([]int, 0)
	for bfs(global.Gragh.Edges, deep, global.Gragh.Head, u, v, &qu) != 0 {
		// 分层后寻找增广路
		minflow := dfs(global.Gragh.Edges, global.Gragh.Head, deep, vis, u, v, inf)
		if minflow != 0 {
			ans += minflow
		}
		for _, item := range qu {
			deep[item] = 0
		}
		qu = make([]int, 0)
	}
	group.Done()git
	return ans
}

func addedge(u int, v int, w int) {
	global.Gragh.Edges = append(global.Gragh.Edges, model.Edge{W: w, From: u, To: v, Next: global.Gragh.Head[u]})
	global.Gragh.Head[u] = global.Gragh.EdgeNumber
	global.Gragh.EdgeNumber++
}

func Match(multiGraph [][]int, n int) {
	Max := 0
	global.Gragh.InitNumberOfNOde = global.Gragh.NodeNumber
	superOriginNode := n
	superConvergeNode := superOriginNode + 1
	OriginNodeList := []int{}
	for i := 0; i < len(multiGraph); i++ {
		// 如果节点是学生+时间，连接上源点
		if len(multiGraph[i]) < 2 {
			continue
		}
		for j := 0; j < len(multiGraph[i]); j++ {
			if multiGraph[i][j] < global.Gragh.NodeNumberOfStu {
				u := multiGraph[i][j]
				//fmt.Printf("%d -> %d||\n",superOriginNode,u)
				addedge(superOriginNode, u, 1)
				addedge(u, superOriginNode, 0)
			} else {
				u := multiGraph[i][j]
				addedge(u, superConvergeNode, 1)
				addedge(superConvergeNode, u, 0)
			}
		}
		Max = int(math.Max(float64(Max), float64(len(multiGraph[i]))))
		OriginNodeList = append(OriginNodeList, superOriginNode)
		global.Gragh.NodeNumber += 2
		superOriginNode += 2
		superConvergeNode += 2
		// 如果节点是学生+第几节课，连接上汇点
	}
	var wg sync.WaitGroup
	for _, item := range OriginNodeList {
		wg.Add(1)
		dinic(item, item+1, &wg)
	}
	wg.Wait()
}

//func MatchPlan2() {
//	superOriginNode := global.Gragh.NodeNumber
//	superConvergeNode := superOriginNode + 1
//	for i := 0; i < global.Gragh.NodeNumber; i++ {
//		if i < global.Gragh.NodeNumberOfStu {
//			addedge(superOriginNode, i, 1)
//			addedge(i, superOriginNode, 0)
//		} else {
//			addedge(i, superConvergeNode, 1)
//			addedge(superOriginNode, i, 0)
//		}
//	}
//	sum := dinic(superOriginNode, superConvergeNode)
//	fmt.Println(sum)
//}

type Pair struct {
	First  int
	Second int
}

func IsMatch(p *Pair) bool {
	stu := strings.Split(global.IndexToStu[p.First], "_")
	te := strings.Split(global.IndexToTe[p.Second], "_")
	i, _ := strconv.ParseInt(stu[1], 0, 0)
	flag := false
	for _, plan := range global.StuToPlan[stu[0]][i].Class {
		if strings.EqualFold(plan, te[1]) {
			flag = true
		}
	}
	if flag {
		for _, teacher := range global.StuToTe[stu[0]] {
			if strings.EqualFold(te[0], strconv.Itoa(int(teacher))) {
				return true
			}
		}
	}
	return false
}

func OutputMatchInfo() []Pair {
	MatchInfo := make([]Pair, 0)
	global.InFirstMatch = make([]bool, global.Gragh.NodeNumber)
	for u := 0; u < global.Gragh.NodeNumberOfStu; u++ {
		for i := global.Gragh.Head[u]; i != -1; i = global.Gragh.Edges[i].Next {
			// 如果边的源点是学生计划节点则不匹配
			// 如果源点是老师放课节点且满流
			if global.Gragh.Edges[i].W == 0 && global.Gragh.Edges[i].From < global.Gragh.NodeNumberOfStu && global.Gragh.Edges[i].To < global.Gragh.InitNumberOfNOde {
				MatchInfo = append(MatchInfo, Pair{global.Gragh.Edges[i].From, global.Gragh.Edges[i].To})
				global.InFirstMatch[global.Gragh.Edges[i].From] = true
				global.InFirstMatch[global.Gragh.Edges[i].To] = true
			}
		}
	}
	return MatchInfo
}
