package kmeans

import (
	"fmt"
	"testing"
)

type TT []int

func (tt TT) TT() {
	tt[0] = 10
}
func TestKmeans(t *testing.T) {

	var mm, mm1, mm2, mm3, mm4 SimpleKeansModel = []float64{3, 4}, []float64{0, 0}, []float64{3, 2}, []float64{5, 8}, []float64{5, 1}
	models := []KeansModel{mm, mm1, mm2, mm3, mm4}
	kmeans := NewKMeans(models, 3)
	kk := kmeans.RandomK()
	clusters, iters := kmeans.Clustering()
	fmt.Println(kk)
	fmt.Println(clusters, iters)
}