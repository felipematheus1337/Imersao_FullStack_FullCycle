package main

import (
	"fmt"
	route2 "github.com/felipematheus1337/Imersao_FullStack_FullCycle/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	lastIndex := len(stringJson) - 1
	fmt.Println(stringJson[lastIndex])
}
