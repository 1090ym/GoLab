package algo

type UF struct {
	id    []int
	sz    []int
	count int
}

func (u *UF) NewUF(N int) {
	for i := 0; i < N; i++ {
		u.id = append(u.id, i)
		u.sz = append(u.sz, 1)
		u.count++
	}
}

func (u *UF) Count() int {
	return u.count
}

func (u *UF) Id() []int {
	return u.id
}

func (u *UF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UF) Union(p int, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	if u.sz[pRoot] > u.sz[qRoot] {
		u.id[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	} else {
		u.id[pRoot] = u.id[qRoot]
		u.sz[qRoot] += u.sz[pRoot]
	}
	u.count--
}
