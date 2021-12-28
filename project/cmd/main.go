package main

import (
	"context"
	"fmt"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
)

func main() {
	f, err := os.Open(".data/ivanovo.osm.pbf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	count := 0
	println(count)
	//686620449 Афанасово 3-я линия 36
	for scanner.Scan() {
		o := scanner.Object()
		switch e := o.(type) {
		case *osm.Node:
			if e.User == "avbruu" {
				fmt.Printf("%f %f \n", e.Lat, e.Lon)

			}
		case *osm.Way:
			if e.User == "avbruu" {
				tags := e.Tags.Map()
				for k, v := range tags {
					fmt.Println(k, v)
				}
			}
		case *osm.Relation:
			if e.User == "avbruu" {
				println("relation")
			}
		}

		continue
	}

	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}
}
