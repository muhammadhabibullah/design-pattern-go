package main

import (
	"fmt"
	"log"

	"design-pattern-go/builder"
)

type director struct {
	builder builder.HouseBuilder
}

func (d *director) setBuilder(builder builder.HouseBuilder) {
	d.builder = builder
}

func (d *director) buildHouse() (builder.House, error) {
	if d.builder == nil {
		return builder.House{}, fmt.Errorf("builder isn't set yet")
	}

	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse(), nil
}

func runBuilder() {
	normalHouseBuilder := builder.GetHouseBuilder(builder.NormalHouseBuilderType)
	iglooHouseBuilder := builder.GetHouseBuilder(builder.IglooHouseBuilderType)

	director := new(director)

	_, err := director.buildHouse()
	if err != nil {
		log.Printf("error build house: %v", err)
	}

	director.setBuilder(normalHouseBuilder)
	normalHouse, _ := director.buildHouse()

	fmt.Printf("\nNormal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.setBuilder(iglooHouseBuilder)
	iglooHouse, _ := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}
