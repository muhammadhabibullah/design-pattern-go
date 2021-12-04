package main

import (
	"flag"
)

func main() {
	factoryMethod := flag.Bool("factory-method", false, "running factory method")
	flag.Parse()

	if *factoryMethod {
		runFactoryMethod()
		return
	}

	flag.PrintDefaults()
}
