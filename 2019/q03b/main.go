package main

import (
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	fileReader := advent2019.FileReader{}
	var grids []*advent2019.Grid

	fileReader.Each(file, func(line string) {
		grid := advent2019.NewGrid()
		grid.Move(line)
		grids = append(grids, grid)
	})

	res := grids[0].QuickestCommonCell(grids[1])
	fmt.Printf("res = %+v\n", res)
}