package generator

type Generator interface {
	SetWindows()
	SetDoorType()
	GenerateHouse() House
}

// GetGenerator returns a generator based on the kind of house
func getGenerator(kind string) Generator {
	switch kind {
	case "ice":
		return NewIce()
	case "wood":
		return NewWood()
	}

	return nil
}
