package main

import (
	"fmt"
	"github.com/avbru/algo/project/utils"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
	tidwal "github.com/tidwall/rtree"
	"net/http"
	"strconv"
	"strings"
)

var result string

func main() {
	p := orb.Point{0, 0}
	_ = project.MercatorScaleFactor(p)
	var tr tidwal.RTree

	nodes, ways, err := utils.ReadPBF(".data/ivanovo.osm.pbf")
	fmt.Println("nodes: ", len(nodes), "polygons: ", len(ways))
	if err != nil {
		panic(err)
	}

	for _, w := range ways {
		if w.Polygon() {
			utils.WayBounds(w, nodes)
			tr.Insert([2]float64{w.Bounds.MinLon, w.Bounds.MinLat}, [2]float64{w.Bounds.MaxLon, w.Bounds.MaxLat}, w)
		}
	}
	tMin, tMax := tr.Bounds()

	fmt.Printf("N: %d tree bounds %f %f \n", tr.Len(), tMin, tMax)
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\n\n"))
		latS, ok := r.URL.Query()["lat"]
		if !ok || len(latS) == 0 {
			w.Write([]byte("no lat \n"))
			return
		}
		lonS, ok := r.URL.Query()["lon"]
		if !ok || len(lonS) == 0 {
			w.Write([]byte("no lon \n"))
			return
		}

		lat, err := strconv.ParseFloat(latS[0], 64)
		if err != nil {
			w.Write([]byte("wrong lat \n"))
			return
		}

		lon, err := strconv.ParseFloat(lonS[0], 64)
		if err != nil {
			w.Write([]byte("wrong lat \n"))
			return
		}
		result = ""
		pt := [2]float64{lon, lat}
		tr.Search(pt, pt,
			func(min, max [2]float64, value interface{}) bool {
				w := value.(*osm.Way)
				result += WayToString(w) + "\n"
				return true
			},
		)

		w.Write([]byte("request: " + lonS[0] + " " + latS[0] + "\n"))
		w.Write([]byte(result))
	}

	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
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
