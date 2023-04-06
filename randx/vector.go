package randx

import "math/rand"

type VectorPlayer struct {
	Index  int
	Vector []int64
}

var _ rand.Source = (*VectorPlayer)(nil)

func NewVectorPlayer(v []int64, index int) *VectorPlayer {
	return &VectorPlayer{
		Vector: v,
		Index:  index,
	}
}

func (p *VectorPlayer) Int63() int64 {
	n := p.Vector[p.Index%len(p.Vector)]
	p.Index++
	return n
}

func (p *VectorPlayer) Seed(seed int64) {
	p.Index = 0
}

func (p *VectorPlayer) Shorten() {
	if p.Index < len(p.Vector) {
		p.Vector = p.Vector[:p.Index]
	}
}

type VectorRecorder struct {
	Rand   *rand.Rand
	Vector []int64
}

var _ rand.Source = (*VectorRecorder)(nil)

func NewVectorRecorder(seed int64) *VectorRecorder {
	vr := &VectorRecorder{}
	vr.Seed(seed)
	return vr
}

func (r *VectorRecorder) Int63() int64 {
	n := r.Rand.Int63()
	r.Vector = append(r.Vector, n)
	return n
}

func (r *VectorRecorder) Seed(seed int64) {
	r.Rand = rand.New(rand.NewSource(seed))
}

func (r *VectorRecorder) Record(n int) []int64 {
	for i := 0; i < n; i++ {
		r.Int63()
	}
	return r.Vector
}
