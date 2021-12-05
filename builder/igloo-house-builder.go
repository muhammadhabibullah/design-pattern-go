package builder

type iglooHouseBuilder struct {
	house House
}

func (b *iglooHouseBuilder) SetWindowType() {
	b.house.WindowType = "Snow Window"
}

func (b *iglooHouseBuilder) SetDoorType() {
	b.house.DoorType = "Snow Door"
}

func (b *iglooHouseBuilder) SetNumFloor() {
	b.house.Floor = 1
}

func (b *iglooHouseBuilder) GetHouse() House {
	return b.house
}
