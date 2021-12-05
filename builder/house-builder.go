package builder

type HouseBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func GetHouseBuilder(houseBuilderType HouseBuilderType) HouseBuilder {
	switch houseBuilderType {
	case NormalHouseBuilderType:
		return &normalHouseBuilder{}
	case IglooHouseBuilderType:
		return &iglooHouseBuilder{}
	default:
		return nil
	}
}
