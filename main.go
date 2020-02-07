package main

import(
	"github.com/common-nighthawk/go-figure"
	"fmt"
)

func main() {
	str := fmt.Stringf("Hello World")
	myFigure := figure.NewFiture(str, "", true)
	  myFigure.Print()
}
