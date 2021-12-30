package utils

import (
	"context"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
)

// ReadPBF - readf osm.pbf file and Project nodes point to Mercator
func ReadPBF(fName string) (map[osm.NodeID]*osm.Node, map[osm.WayID]*osm.Way, error) {
	f, err := os.Open(fName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	//686620449 Афанасово 3-я линия 36
	nodes := make(map[osm.NodeID]*osm.Node, 4_000_000)
	ways := make(map[osm.WayID]*osm.Way, 400_000)
	for scanner.Scan() {
		o := scanner.Object()
		switch e := o.(type) {
		case *osm.Node:
			//Cast to cartesian
			p := orb.Point{e.Lon, e.Lat}
			p = project.Point(p, project.WGS84.ToMercator)
			e.Lon, e.Lat = p.Lon(), p.Lat()
			nodes[e.ID] = e
		case *osm.Way:
			ways[e.ID] = e
		}

		scanErr := scanner.Err()
		if scanErr != nil {
			return nil, nil, scanErr
		}
	}
	return nodes, ways, nil
}
