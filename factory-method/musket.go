package factoryMethod

type musket struct {
	gun
}

func newMusket() Gun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
