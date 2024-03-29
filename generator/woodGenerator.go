package generator

type wood struct {
	windowType string
	doorType   string
}

func NewWood() *wood {
	return &wood{}
}

func (w *wood) SetWindows() {
	w.windowType = "wood"
}

func (w *wood) SetDoorType() {
	w.doorType = "wood.door"
}

// GenerateHouse returns a house based on the wood generator
func (w *wood) GenerateHouse() House {
	return House{
		WindowType: w.windowType,
		DoorType:   w.doorType,
	}
}
