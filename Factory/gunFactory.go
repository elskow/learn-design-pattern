package Factory

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
