package bridge

import "fmt"

type printer interface {
	print() string
}

type hp struct {
}

func (h *hp) print() string {
	return fmt.Sprintf("this is hp printer")
}

type sony struct {
}

func (s *sony) print() string {
	return fmt.Sprintf("this is sony printer")
}
