package main

import (
	"fmt"

	route "github.com/EdmilsonFerreiraF/positions-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
}
