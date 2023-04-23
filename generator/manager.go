package generator

type manage struct {
	generator Generator
}

func NewManage(generator Generator) *manage {
	return &manage{generator: generator}
}

func (m *manage) BuildHouse() House {
	m.generator.SetWindows()
	m.generator.SetDoorType()

	return m.generator.GenerateHouse()
}
