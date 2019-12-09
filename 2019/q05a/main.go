package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	in := make(chan big.Int, 1)
	out := make(chan big.Int, 10)

	in <- *big.NewInt(1)

	c := advent2019.NewComputer(in, out)
	c.SetInput(strings.TrimSpace(string(all)))
	c.Calculate()
	close(out)

	for res := range out {
		fmt.Printf("res = %+v\n", res.String())
	}
}
