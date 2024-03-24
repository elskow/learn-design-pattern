package AbstractFactory

type WesternGoMakanFactory struct{}

func (WesternGoMakanFactory) CreateMakanan() IMakanan {
	return &MakananWestern{
		Makanan: Makanan{
			nama:  "Burger",
			price: 20000,
		},
	}
}

func (WesternGoMakanFactory) CreateMinuman() IMinuman {
	return &MinumanWestern{
		Minuman: Minuman{
			nama:    "Coca Cola",
			price:   10000,
			kondisi: Dingin,
		},
	}
}
