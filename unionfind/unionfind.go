package unionfind

type UnionFind struct {
	parent []int
}

func NewUnionFind(N int) *UnionFind {
	set := []int{}
	for i := 0; i < N; i++ {
		set = append(set, i)
	}
	return &UnionFind{set}
}

func (u *UnionFind) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])

	}
	return u.parent[x]
}
func (u *UnionFind) Union(x int, y int) {
	u.parent[u.Find(x)] = u.Find(y)
}
