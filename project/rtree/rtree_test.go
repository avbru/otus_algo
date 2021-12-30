package rtree

import "testing"

var entries = map[string]*Entry{
	"L1": {
		Low:   [2]float64{1.0, 2.0},
		High:  [2]float64{1.0, 2.0},
		Value: "L1",
	},
	"LM": {
		Low:   [2]float64{3.0, 4.0},
		High:  [2]float64{4.0, 5.0},
		Value: "L1",
	},
}

func TestRTree(t *testing.T) {
	tr := New(2, 1)
	tr.Insert(entries["L1"])
	tr.Insert(entries["LM"])

	tr.String()
}
