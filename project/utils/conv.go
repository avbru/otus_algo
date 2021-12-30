package utils

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/osm"
)

// WayBounds - compute way bound rectangle
func WayBounds(w *osm.Way, nodes map[osm.NodeID]*osm.Node) {
	if w.Polygon() {

		ring := orb.Ring{}
		for _, n := range w.Nodes {
			nd, ok := nodes[n.ID]
			if ok {
				ring = append(ring, orb.Point{nd.Lon, nd.Lat})
			}
		}
		b := ring.Bound()
		w.Bounds = &osm.Bounds{
			MinLat: b.Min.Lat(),
			MaxLat: b.Max.Lat(),
			MinLon: b.Min.Lon(),
			MaxLon: b.Max.Lon(),
		}
	}
}
