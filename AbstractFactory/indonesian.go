package AbstractFactory

type IndonesianGoMakanFactory struct{}

func (i *IndonesianGoMakanFactory) CreateMakanan() IMakanan {
	return &MakananIndonesian{
		Makanan: Makanan{
			nama:  "Nasi Goreng",
			price: 15000,
		},
	}
}

func (i *IndonesianGoMakanFactory) CreateMinuman() IMinuman {
	return &MinumanIndonesian{
		Minuman: Minuman{
			nama:  "Es Teh",
			price: 5000,
		},
	}
}
