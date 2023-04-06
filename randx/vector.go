package randx

import "math/rand"

type Vector []int64

type VectorPlayer struct {
	Index  int
	Vector Vector
}

var _ rand.Source = (*VectorPlayer)(nil)

func NewVectorPlayer(v Vector, index int) *VectorPlayer {
	return &VectorPlayer{
		Vector: v,
		Index:  index,
	}
}

func (v VectorPlayer) Int63() int64 {
	return v.Vector[v.Index%len(v.Vector)]
}

func (v *VectorPlayer) Seed(seed int64) {
}

type VectorRecorder struct {
	Rand   *rand.Rand
	Vector Vector
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

type ProvisionedVectorRecorder struct {
	*VectorRecorder
	Index int
}

var _ rand.Source = (*ProvisionedVectorRecorder)(nil)

func NewProvisionedVectorRecorder(seed int64, maxSize int, index int64) *ProvisionedVectorRecorder {
	r := &ProvisionedVectorRecorder{
		VectorRecorder: NewVectorRecorder(seed),
		Index:          int(index),
	}

	for i := 0; i < maxSize; i++ {
		r.Int63()
	}

	return r
}

func (r *ProvisionedVectorRecorder) Int63() int64 {
	n := r.Vector[r.Index%len(r.Vector)]
	r.Index++
	return n
}

func (r *ProvisionedVectorRecorder) Seed(seed int64) {
	r.VectorRecorder.Seed(seed)
}
