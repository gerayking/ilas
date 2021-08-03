package main

import (
	"container/list"
)

// dfs分割
func dfs_to(g *graph, v int, sonGraph list.List, vis [NODENUMBER]bool) {
	sonGraph.PushBack(v)
	for i := g.head[v]; i != -1; i = g.edges[i].next {
		t := g.edges[i].to
		if !vis[t] {
			dfs_to(g, t, sonGraph, vis)
		}
		sonGraph.Init()
	}
}

// 进行dfs将图进行分割
func tarjan(g *graph, n int) list.List {
	vis := [NODENUMBER]bool{}
	graphdivident := list.List{}
	sgraph := list.List{}
	for i := 0; i < n; i++ {
		if vis[i] == false {
			dfs_to(g, i, sgraph, vis)
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
func bfs(edges *[NODENUMBER]edge, deep *[NODENUMBER]int, head *[NODENUMBER]int, s int, t int) int {
	queue := New()
	queue.Push(0)
	for queue.length != 0 {
		queue.Pop()
	}
	deep[s] = 1
	queue.Push(s)
	for queue.length != 0 {
		u := queue.top.value.(int)
		queue.Pop()
		i := 0
		for i != -1 {
			i = head[u]
			if edges[i].w > 0 && deep[edges[i].to] == 0 {
				deep[edges[i].to] = deep[edges[i].from] + 1
				queue.Push(edges[i].to)
			}
			i = edges[i].next
		}
	}
	if deep[t] == 0 {
		return 0
	}
	return 1
}

// 寻找增广路
func dfs(edges *[NODENUMBER]edge, head *[NODENUMBER]int, deep *[NODENUMBER]int, u int, t int, dist int) int {
	if u == t {
		return dist
	}
	i := 0
	for i != -1 {
		i = head[u]
		if deep[edges[i].to] == deep[edges[i].from]+1 && edges[i].w != 0 {
			di := dfs(edges, head, deep, edges[i].to, t, dist)
			if di > 0 {
				edges[i].w -= di
				edges[i^1].w += di
				return di
			}
		}
	}
	return 0
}

// dinic 计算最大流
func dinic(g graph, u int, v int) int {
	ans := 0
	var deep [NODENUMBER]int
	// 对残流图不断进行分层
	for bfs(&g.edges, &deep, &g.head, u, v) != 0 {
		// 分层后寻找增广路
		minflow := dfs(&g.edges, &g.head, &deep, u, v, inf)
		if minflow != 0 {
			ans += minflow
		}
	}
	return ans
}
func addedge(u int, v int, g graph) {
	g.edges[g.edgeNumber].w = 1
	g.edges[g.edgeNumber].from = u
	g.edges[g.edgeNumber].to = v
	g.edges[g.edgeNumber].next = g.head[u]
	u++
}
func Match(g graph, multiGraph list.List, n int) {
	superOriginNode := n
	superConvergeNode := superOriginNode + 1
	OriginNodeList := list.List{}
	for item := multiGraph.Front(); item != nil; item = item.Next() {
		// 如果节点是老师+时间，连接上源点
		if true {
			addedge(superOriginNode, item.Value.(int), g)
		} else {
			addedge(item.Value.(int), superConvergeNode, g)
		}
		OriginNodeList.PushFront(superOriginNode)
		superOriginNode += 2
		superConvergeNode += 2
		// 如果节点是学生+第几节课，连接上汇点
	}
	for item := OriginNodeList.Front(); item != nil; item = item.Next() {
		superOriginNode := item.Value.(int)
		superConvergeNode := superOriginNode + 1
		dinic(g, superOriginNode, superConvergeNode)
	}
}

type Pair struct {
	first  int
	second int
}

func outputMatchInfo(g graph) list.List {
	MatchInfo := list.List{}
	for u := 0; u < g.nodeNumber; u++ {
		// 如果边的源点是学生计划节点则不匹配
		// 坑坑坑坑坑坑坑坑坑坑坑坑坑
		for v := g.head[u]; v != -1; v = g.edges[v].next {
			// 如果边的源点是学生计划节点则不匹配
			// 如果源点是老师放课节点且满流
			// 坑坑坑坑坑坑坑坑坑坑坑坑坑
			if g.edges[v].w == 0 {
				MatchInfo.PushFront(Pair{u, v})
			}
		}
	}
	return MatchInfo
}
