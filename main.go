package main

import (
	"flag"
)

func main() {
	factoryMethod := flag.Bool("factory-method", false, "running factory method")
	builder := flag.Bool("builder", false, "running builder")
	flag.Parse()

	if *factoryMethod {
		runFactoryMethod()
		return
	}
	if *builder {
		runBuilder()
		return
	}

	flag.PrintDefaults()
}
