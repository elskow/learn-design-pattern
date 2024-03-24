package AbstractFactory

type IGoMakanFactory interface {
	CreateMakanan() IMakanan
	CreateMinuman() IMinuman
}

func GetGoMakanFactory(factoryType string) IGoMakanFactory {
	if factoryType == "Indonesian" {
		return &IndonesianGoMakanFactory{}
	}
	if factoryType == "Western" {
		return WesternGoMakanFactory{}
	}
	return nil
}
