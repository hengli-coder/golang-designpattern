package bridge

import (
	"fmt"
	"testing"
)

func TestMac(t *testing.T) {
	m := &mac{name: "mac"}
	w := &windows{name: "windows"}
	s := &sony{}
	h := &hp{}

	fmt.Println("this is an example")
	m.setPrinter(s)
	fmt.Println(m.print())

	m.setPrinter(h)
	fmt.Println(m.print())

	w.setPrinter(s)
	fmt.Println(w.print())
	w.setPrinter(h)
	fmt.Println(w.print())
	//mac: this is sony printer
	//mac: this is hp printer
	//windows: this is sony printer
	//windows: this is hp printer
}
