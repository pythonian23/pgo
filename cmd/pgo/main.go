package main

import (
	"fmt"
	"os"

	"github.com/pythonian23/pgo"
)

func main() {
	fmt.Println(pgo.Parse(os.Args))
}
