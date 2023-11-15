package heap

type Present struct {
	Value int
	Size  int
}

type Presents []*Present

func (p Presents) Len() int { return len(p) }

func (p Presents) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Size < p[j].Size
	}

	return p[i].Value > p[j].Value
}

func (p Presents) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Presents) Push(x interface{}) {
	*p = append(*p, x.(*Present))
}

func (p *Presents) Pop() interface{} {
	old := *p
	n := old.Len()
	x := old[n-1]
	*p = old[:n-1]
	return x
}
