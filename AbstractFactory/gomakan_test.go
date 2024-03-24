package AbstractFactory

import "testing"

func TestGoMakanFactory(t *testing.T) {
	westernFactory := GetGoMakanFactory("Western")
	indonesianFactory := GetGoMakanFactory("Indonesian")

	if westernFactory == nil {
		t.Error("Expected WesternGoMakanFactory, but got nil")
	}
	if indonesianFactory == nil {
		t.Error("Expected IndonesianGoMakanFactory, but got nil")
	}

	westernMakanan := westernFactory.CreateMakanan()
	indonesianMakanan := indonesianFactory.CreateMakanan()

	if westernMakanan.getNama() != "Burger" {
		t.Errorf("Expected Burger, but got %s", westernMakanan.getNama())
	}

	if indonesianMakanan.getNama() != "Nasi Goreng" {
		t.Errorf("Expected Nasi Goreng, but got %s", indonesianMakanan.getNama())
	}
}
