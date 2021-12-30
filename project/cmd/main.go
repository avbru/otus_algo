package main

import (
	"fmt"
	"github.com/avbru/algo/project/utils"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
	"github.com/tidwall/rtree"
	"strings"
)

func main() {
	p := orb.Point{0, 0}
	_ = project.MercatorScaleFactor(p)
	var tr rtree.RTree

	nodes, ways, err := utils.ReadPBF(".data/ivanovo.osm.pbf")
	fmt.Println("nodes: ", len(nodes))
	fmt.Println("ways: ", len(ways))
	if err != nil {
		panic(err)
	}

	for _, w := range ways {
		//for key, val := range w.TagMap() {
		//if strings.Contains(val, "Афанасово") {

		//	println(key, val, w.Tags.Find("addr:street"), w.Tags.Find("addr:housenumber"), "N-nodes:", len(w.Nodes))

		if w.Polygon() {

			utils.WayBounds(w, nodes)
			//fmt.Printf("min.X %f min.Y %f\nmax.X %f max.Y %f\n",
			//	w.Bounds.MinLon, w.Bounds.MinLat,
			//	w.Bounds.MaxLon, w.Bounds.MaxLat)
			//	fmt.Printf("%f\n", (w.Bounds.MaxLat-w.Bounds.MinLat)*(w.Bounds.MaxLon-w.Bounds.MinLon))

			tr.Insert([2]float64{w.Bounds.MinLon, w.Bounds.MinLat}, [2]float64{w.Bounds.MaxLon, w.Bounds.MaxLat}, w)
		}
		//}
		//}

	}
	//Lat Lon
	//57.0196868 41.0370372
	//pt := [2]float64{41.0370372, 57.0196868}

	pt := [2]float64{4568240, 7764167} //Афанасово ул. Успенская д.11
	//pt := [2]float64{4565167.595445, 7758821.770429} //Суворова 72
	tMin, tMax := tr.Bounds()

	fmt.Printf("N: %d tree bounds %f %f \n", tr.Len(), tMin, tMax)
	tr.Search(pt, pt,
		func(min, max [2]float64, value interface{}) bool {
			w := value.(*osm.Way)
			println(WayToString(w))
			return true
		},
	)
}

func WayToString(w *osm.Way) string {
	tags := w.TagMap()
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("objectID: %v\n", w.ID))
	for k, v := range tags {
		s.WriteString(fmt.Sprintf("%s %s\n", k, v))
	}
	return s.String()
}
