package main

import (
	"fmt"
	"log"

	"design-pattern-go/factory-method"
)

func runFactoryMethod() {
	ak47, _ := factoryMethod.GetGun(factoryMethod.AK47GunType)
	fmt.Println(ak47.GetName())
	fmt.Println(ak47.GetPower())

	musket, _ := factoryMethod.GetGun(factoryMethod.MusketGunType)
	fmt.Println(musket.GetName())
	fmt.Println(musket.GetPower())

	_, err := factoryMethod.GetGun("toy")
	if err != nil {
		log.Printf("error get gun: %v", err)
	}
}
