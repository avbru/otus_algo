[дампы по регионам](http://osm.sbin.ru/osm_dump/)

Mercator карта мира http://epsg.io/map#srs=3857&x=4568222&y=7764143&z=19&layer=streets

Успенская д.11 bound
Lat         Lon
57.0196868 41.0370372
57.0199197 41.0373869


Меркатор  Y-77**** X=45****
7764143.548870685 4568222.084768625
7764191.176960677 4568261.013194555


fmt.Printf("min.X %f min.Y %f\nmax.X %f max.Y %f\n",
w.Bounds.MinLon, w.Bounds.MinLat,
w.Bounds.MaxLon, w.Bounds.MaxLat)

minX 4568222.084769 minY 7764143.548871
max.X 4568261.013195 max.Y 7764191.176961
