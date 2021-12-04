package factoryMethod

import (
	"fmt"
)

func GetGun(gunType GunType) (Gun, error) {
	switch gunType {
	case AK47GunType:
		return newAk47(), nil
	case MusketGunType:
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("invalid gunType: %s", gunType)
	}
}
