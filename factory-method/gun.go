package factoryMethod

type Gun interface {
	SetName(name string)
	SetPower(power int64)
	GetName() string
	GetPower() int64
}

type gun struct {
	name  string
	power int64
}

func (g *gun) SetName(name string) {
	g.name = name
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(power int64) {
	g.power = power
}

func (g *gun) GetPower() int64 {
	return g.power
}
