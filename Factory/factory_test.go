package Factory

import "testing"

func TestGetGun(t *testing.T) {
	ak47, _ := getGun("ak47")
	if ak47.getName() != "AK47 gun" {
		t.Error("Expected AK47 gun but got ", ak47.getName())
	}
	if ak47.getPower() != 4 {
		t.Error("Expected power 4 but got ", ak47.getPower())
	}

	musket, _ := getGun("musket")
	if musket.getName() != "Musket gun" {
		t.Error("Expected Musket gun but got ", musket.getName())
	}
	if musket.getPower() != 1 {
		t.Error("Expected power 1 but got ", musket.getPower())
	}
}
