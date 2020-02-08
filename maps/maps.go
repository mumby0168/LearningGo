package main

import "fmt"

type GeoPoint struct {
	Lat, Long float64
}

func simpleMap() {
	favPlaces := make(map[string]GeoPoint)
	favPlaces["My House"] = GeoPoint{
		24.6666,
		52.4354,
	}

	fmt.Println(favPlaces["My House"])
}

func mapOperations() {
	favPlaces := make(map[string]GeoPoint)
	favPlaces["My House"] = GeoPoint{
		24.6666,
		52.4354,
	}

	_, ok := favPlaces["Work"]
	fmt.Println(ok)

	delete(favPlaces, "My House")
	fmt.Println(favPlaces)

}

func main() {
	simpleMap()
	mapOperations()
}
