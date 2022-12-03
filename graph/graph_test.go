package graph

import (
	"fmt"
	"testing"
)

var output = func(i int) {
	fmt.Printf("->%v", i)
}

func TestNewGraph(t *testing.T) {
	ug := NewUGraph(10)
	/*
				1--2--4
				|   \
				3    5
		        |
		        6
	*/
	ug.AddEdge(1, 2)
	ug.AddEdge(1, 3)
	ug.AddEdge(2, 4)
	ug.AddEdge(2, 5)
	ug.AddEdge(3, 6)
	fmt.Println("\n----DFS------")
	// 1 2 4 5 3 6
	ug.DFS(1, output)
	fmt.Println("\n---DFS-Iterate-------")
	ug.DFS_Iterate(1, output)
	fmt.Println("\n---BFS-------")
	ug.BFS(1, output)
	fmt.Println("")
}

func TestJudgeCircle(t *testing.T) {
	ug := NewUGraph(7)
	// ug.AddEdge(0, 1)
	// ug.DFS(0, output)
	ug.AddEdge(0, 1).AddEdge(0, 2).AddEdge(0, 3).AddEdge(1, 3)
	fmt.Println(ug.IsCyclic(0), ug.IsCyclic2(0)) //true
	ug.RmEdge(1, 3)
	fmt.Println(ug.IsCyclic(0), ug.IsCyclic2(0)) //false
	//

}

func TestXxx(t *testing.T) {
	dg := NewGraph(4, true)
	dg.AddEdge(1, 2)
	dg.AddEdge(2, 3)
	dg.AddEdge(3, 1)
	/*
			     1
		//      / ^
		//     v   \
			  2  -> 3
	*/
	fmt.Println("判断有向图dg是否为环:", dg.IsCyclic(1))
	dg.RmEdge(3, 1)
	fmt.Println("判断有向图dg是否为环:", dg.IsCyclic(1))
	dg.AddEdge(1, 3)
	fmt.Println("判断有向图dg是否为环:", dg.IsCyclic(1))
}
