package builder

type normalHouseBuilder struct {
	house House
}

func (b *normalHouseBuilder) SetWindowType() {
	b.house.WindowType = "Wooden Window"
}

func (b *normalHouseBuilder) SetDoorType() {
	b.house.DoorType = "Wooden Door"
}

func (b *normalHouseBuilder) SetNumFloor() {
	b.house.Floor = 2
}

func (b *normalHouseBuilder) GetHouse() House {
	return b.house
}
