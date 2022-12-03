package graph

import (
	"fmt"
	"kaoyan/queue"
	"kaoyan/stack"
	"kaoyan/unionfind"
)

type Graph struct {
	Vertex     [][]int
	record     []int
	size       int
	isDirected bool
}

func NewGraph(size int, dir bool) *Graph {
	vertex := make([][]int, 0)
	for i := 0; i < size; i++ {
		vertex = append(vertex, make([]int, size))
	}
	// fmt.Printf("%v", vertex)
	return &Graph{vertex, make([]int, size), size, dir}
}

func (g *Graph) Print() {
	painter := ""
	for _, v := range g.Vertex {
		for _, vv := range v {
			painter += fmt.Sprint(vv) + ","
		}
		painter += "\n"

	}
	fmt.Print(painter)
}

func (g *Graph) AddEdge(from, to int) {
	g.Vertex[from][to] = 1
}
func (g *Graph) RmEdge(from, to int) {
	g.Vertex[from][to] = 0
}
func (g *Graph) ResetRecord() {
	g.record = make([]int, g.size)
}

func (g *Graph) dfs(v int, visit func(i int)) {
	if g.record[v] == 1 {
		return
	}
	// 访问记录
	visit(v)
	g.record[v] = 1

	// 对邻居结点递归dfs
	for i := 0; i < g.size; i++ {
		if g.Vertex[v][i] != 0 {
			if i == v && g.record[v] == 1 { //表示该结点不是刚才对那个结点的情况下, 已经被访问了
				fmt.Println("不是环！")
			}
			g.dfs(i, visit)
		}

	}
}
func (g *Graph) dfs_iter2(v int, visit func(int)) {
	stack := stack.NewStack[int]()
	/* 单独把该点遍历，然后入队 */
	stack.Push(v)
	g.record[v] = 1
	for !stack.IsEmpty() {
		j := stack.Pop()
		visit(j)
		for _, vv := range g.getNeighbour(j) {
			if g.record[vv] == 0 {
				stack.Push(vv)
				g.record[vv] = 1
			}
		}
	}
}
func (g *Graph) dfs_iter(v int, visit func(int)) {
	stack := stack.NewStack[int]()
	/* 单独把该点遍历，然后入栈 */

	g.record[v] = 1
	stack.Push(v)
	for !stack.IsEmpty() {
		j := stack.Pop()
		visit(j)
		for i := g.size - 1; i >= 0; i-- {
			if g.Vertex[j][i] != 0 && g.record[i] == 0 { //如果相邻且没有被访问
				g.record[i] = 1
				stack.Push(i)
			}
		}
	}
}
func (g *Graph) bfs(v int, visit func(int)) {
	queue := queue.NewQueue(len(g.Vertex[v]))
	/* 单独把该点遍历，然后入队 */
	g.record[v] = 1
	queue.Enqueue(v)
	for !queue.IsEmpty() {
		j := queue.Dequeue()
		visit(j)
		for _, i := range g.getNeighbour(j) {
			if g.record[i] == 0 { //如果相邻且没有被访问
				g.record[i] = 1
				queue.Enqueue(i)
			}
		}
	}
}
func (g *Graph) getNeighbour(cur int) []int {
	neis := []int{}
	for i := 0; i < g.size; i++ {
		if g.Vertex[cur][i] == 1 {
			neis = append(neis, i)
		}
	}
	return neis
}

/* 判断是非是环 */
func (g *Graph) IsCyclicHelper(cur int, visited []bool, par *int) bool {
	visited[cur] = true
	for _, nei := range g.getNeighbour(cur) {
		if par != nil {
			fmt.Printf("%v-%v\n", nei, *par)
		}

		if !visited[nei] {
			if g.IsCyclicHelper(nei, visited, &cur) {
				return true
			}
		} else {
			if par != nil && nei != *par {
				return true
			}
		}
	}
	return false
}
func (g *Graph) IsCyclic(start int) bool {
	visited := make([]bool, g.size)
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			if g.IsCyclicHelper(i, visited, nil) {
				return true
			}
		}
	}
	return false

}

/* BFS和DFS */
func (g *Graph) DFS(v int, visit func(int)) {
	g.dfs(v, visit)
	g.ResetRecord()
}
func (g *Graph) DFS_Iterate(v int, visit func(int)) {
	g.dfs_iter(v, visit)
	g.ResetRecord()
}
func (g *Graph) BFS(v int, visit func(int)) {
	g.bfs(v, visit)
	g.ResetRecord()

}

/* 无向图 */
type UndirectedGraph struct{ Graph } //基于组合概念，无向图是普通图的一个组合
func NewUGraph(size int) *UndirectedGraph {

	return &UndirectedGraph{*NewGraph(size, false)}
}
func (ug *UndirectedGraph) AddEdge(a, b int) *UndirectedGraph {
	ug.Graph.AddEdge(a, b)
	ug.Graph.AddEdge(b, a)
	return ug
}

func (ug *UndirectedGraph) RmEdge(a, b int) {
	ug.Graph.RmEdge(a, b)
	ug.Graph.RmEdge(b, a)
}

/* 判断环(只能用于无向图)--UnionFind法 */
func (ug *UndirectedGraph) IsCyclic2(start int) bool {
	type EdgesSet map[[2]int]int

	edges := make(EdgesSet)
	u := unionfind.NewUnionFind(ug.size * ug.size)

	for i := 0; i < ug.size; i++ {
		for _, v := range ug.getNeighbour(i) {
			list := [2]int{}
			if v >= i {
				list[0] = v
				list[1] = i
			} else {
				list[1] = v
				list[0] = i
			}
			edges[list] = 1
			// fmt.Println("edges", edges)
		}
	}

	for edge := range edges {
		if u.Find(edge[0]) == u.Find(edge[1]) {
			return true
		}
		u.Union(edge[1], edge[0])
	}
	return false
}
