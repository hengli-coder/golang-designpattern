package generator

import "fmt"

func build() {
	ice := getGenerator("ice")
	m := NewManage(ice)
	house := m.BuildHouse()

	fmt.Println(house)

	wood := getGenerator("wood")
	m = NewManage(wood)
	house = m.BuildHouse()
	fmt.Println(house)
}
