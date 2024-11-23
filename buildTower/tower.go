package main

import (
	"fmt"
	"strings"
)

func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	nAsteriks := 2*nFloors - 1
	nSpace := 0
	for i := nFloors; i > 0; i-- {
		space := strings.Repeat(" ",nSpace)
		blocks := strings.Repeat("*",nAsteriks)
		tower[i-1] = space+blocks+blocks

		nAsteriks -= 2
		nSpace += 1
	}

	return tower
}

func main() {
	fmt.Println(TowerBuilder(6))
}