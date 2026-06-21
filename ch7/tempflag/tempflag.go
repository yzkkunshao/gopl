// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/linehk/gopl/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	fmt.Printf("%q", os.Args)
	flag.Parse()
	fmt.Println(*temp)
}
