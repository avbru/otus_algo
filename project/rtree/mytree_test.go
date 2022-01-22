package rtree

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_PeekSeeds(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	//         0  1  2  3
	//var f, s []int //arrays
	h := len(a) / 2
	fmt.Println(a[:h], a[h:])
	for i := 0; i < h; i++ { //Cycle over positions
		for j := h; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			fmt.Println(a[:h], a[h:])
			a[i], a[j] = a[j], a[i]
		}
	}
}

func Test_Perm(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	//         0  1  2  3
	//var f, s []int //arrays
	h := len(a) / 2
	fmt.Println(a[:h], a[h:])
	for i := 0; i < h; i++ { //Cycle over positions
		for j := h; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			fmt.Println(a[:h], a[h:])
			a[i], a[j] = a[j], a[i]
		}
	}
}

type entry struct {
	rec  Rec
	data string
}

func TestRtree_Insert(t *testing.T) {
	es := make([]*entry, 0)
	es = append(es, &entry{rec: Rec{Low: Point{11, 7}, Top: Point{11.5, 15}}, data: "R13"})
	es = append(es, &entry{rec: Rec{Low: Point{10, 11}, Top: Point{11.25, 12.5}}, data: "R14"})

	es = append(es, &entry{rec: Rec{Low: Point{1, 1}, Top: Point{3, 3}}, data: "R15"})
	es = append(es, &entry{rec: Rec{Low: Point{4, 2}, Top: Point{7, 5}}, data: "R16"})
	es = append(es, &entry{rec: Rec{Low: Point{8, 4}, Top: Point{9, 16}}, data: "R11"})
	es = append(es, &entry{rec: Rec{Low: Point{5.5, 6}, Top: Point{8.5, 8}}, data: "R12"})
	es = append(es, &entry{rec: Rec{Low: Point{6, 10}, Top: Point{7.5, 11.5}}, data: "R10"})
	es = append(es, &entry{rec: Rec{Low: Point{6, 13}, Top: Point{7.5, 14.5}}, data: "R9"})
	es = append(es, &entry{rec: Rec{Low: Point{3, 9}, Top: Point{4.5, 11.5}}, data: "R8"})
	//
	es = append(es, &entry{rec: Rec{Low: Point{14, 2}, Top: Point{15, 7}}, data: "R18"})
	es = append(es, &entry{rec: Rec{Low: Point{13, 4}, Top: Point{17, 6}}, data: "R17"})
	es = append(es, &entry{rec: Rec{Low: Point{14.5, 3}, Top: Point{16, 5}}, data: "R19"})
	//
	//es = append(es, &entry{rec: Rec{Low: Point{11, 7}, Top: Point{11.5, 15}}, data: "R13"})
	//es = append(es, &entry{rec: Rec{Low: Point{10, 11}, Top: Point{11.25, 12.5}}, data: "R14"})
	tr := New(10)
	for _, v := range es {
		tr.Insert(v.rec, v.data)
	}
	tr.Draw()

	println("-----------")
	tr.Search(Rec{Low: Point{1, 1}, Top: Point{1, 1}})
	println("-----------")
	tr.Search(Rec{Low: Point{14.6, 3}, Top: Point{14.6, 4.8}})
}

func Test_bbox(t *testing.T) {
	r1 := Rec{
		Low: Point{1.0, 1.0},
		Top: Point{2.0, 2.0},
	}

	r2 := Rec{
		Low: Point{4.0, 1.0},
		Top: Point{5.0, 2.0},
	}
	require.Equal(t, 4.0, bbox(r1, r2).S())
}

func Test_RecS(t *testing.T) {
	r := Rec{
		Low: Point{-2.0, -2.0},
		Top: Point{2.0, 2.0},
	}
	require.Equal(t, 16.0, r.S())
}
