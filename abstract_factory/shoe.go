package abstract_factory

type iShoe interface {
	getSize() int
	getPrice() float64
}

type shoe struct {
	size  int
	color string
	price float64
	style string
}

func (s *shoe) getSize() int {
	return s.size
}

func (s *shoe) getPrice() float64 {
	return s.price
}
