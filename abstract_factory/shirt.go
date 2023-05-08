package abstract_factory

type iShirt interface {
	getColor() string
	getSize() int
}

type shirt struct {
	size     int
	color    string
	price    float64
	material string
}

func (s *shirt) getColor() string {
	return s.color
}

func (s *shirt) getSize() int {
	return s.size
}
