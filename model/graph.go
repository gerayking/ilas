package model

//type studentNode struct {
//	nodeId    int
//	studentId int
//	timeId    int
//}
//type teacherNode struct {
//	nodeId    int
//	teacherId int
//	timeId    int
//}
type Edge struct {
	From int // 图的源点
	To   int // 图的终点
	Next int // 下一条边
	W    int // 边的流量
}

type Graph struct {
	Head       []int  // 邻接表的头节点
	Edges      []Edge // 邻接表表示边
	EdgeNumber int    // 边的数量
	NodeNumber int    // 点的数量
}
