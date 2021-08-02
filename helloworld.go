package main

const N = 10000000
const inf = 9999999999

type node struct {
	from  int
	to    int
	next  int
	w     int
	value int
}

func buildgraph() {
	return
}
func bfs(edges [N]node, deep [N]int, head [N]int, s int, t int) int {
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
func dfs(edges [N]node, head [N]int, deep [N]int, u int, t int, dist int) int {
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
func dinic(edges [N]node, deep [N]int, head [N]int, u int, v int) int {
	ans := 0
	for bfs(edges, deep, head, u, v) != 0 {
		minflow := dfs(edges, head, deep, u, v, inf)
		if minflow != 0 {
			ans += minflow
		}
	}
	return ans
}
func main() {
	head := [N]int{}
	edge := [N]node{}
	deep := [N]int{}
	buildgraph()
	dinic(edge, deep, head, 0, 1000)
}
