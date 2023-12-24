package concate

import "testing"

const (
	concateStep = 1000
)

func BenchmarkPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		simplePlus(concateStep)
	}
}

func BenchmarkBuff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf(concateStep)
	}
}

func BenchmarkBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		builder(concateStep)
	}
}
