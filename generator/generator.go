package generator

type Generator interface {
	SetWindows()
	SetDoorType()
	GenerateHouse() House
}

func getGenerator(kind string) Generator {
	switch kind {
	case "ice":
		return NewIce()
	case "wood":
		return NewWood()
	}

	return nil
}
