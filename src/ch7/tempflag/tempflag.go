package main

import(
	"flag"
	"fmt"

	"ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp",20.0,"the temperature")

func main(){
	flag.Parse()
	fmt.Println(*temp)
}
