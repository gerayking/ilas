package model

type Edge struct {
	From int // 图的源点
	To   int // 图的终点
	Next int // 下一条边
	W    int // 边的流量
}

type Graph struct {
	Head             []int  // 邻接表的头节点
	Edges            []Edge // 邻接表表示边
	EdgeNumber       int    // 边的数量
	NodeNumber       int    // 点的数量
	NodeNumberOfStu  int    // 左部图和右部图的边界
	InitNumberOfNOde int    // 未添加超级源点和超级汇点前的节点数量
}

// 临时节点，用于重新建图
type TmpNode struct {
	TeacherId string
	LessonId  []int
}

// 用于存储匹配信息
type Pair struct {
	First  int
	Second int
}
