package generator

type ice struct {
	windowType string
	doorType   string
	price      int
}

func NewIce() *ice {
	return &ice{
		price: 1000,
	}
}

func (i *ice) SetWindows() {
	i.windowType = "ice.window"
}

func (i *ice) SetDoorType() {
	i.doorType = "ice.door"
}

func (i *ice) GenerateHouse() House {
	return House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Price:      i.price,
	}
}
