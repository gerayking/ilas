package main

const NODENUMBER = 10000000
const EDGESNUMBER = 10000000

type studentNode struct {
	nodeId    int
	studentId int
	timeId    int
}
type edge struct {
	from int // 图的源点
	to   int // 图的终点
	next int // 下一条边
	w    int // 边的流量
}
type teacherNode struct {
	nodeId    int
	teacherId int
	timeId    int
}
type graph struct {
	head       [NODENUMBER]int   // 邻接表的头节点
	edges      [EDGESNUMBER]edge // 邻接表表示边
	edgeNumber int               // 边的数量
	nodeNumber int               // 点的数量
}
