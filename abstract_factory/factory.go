package abstract_factory

type SportsFactory interface {
	makeShoes() iShoe
	makeShirts() iShirt
}

func createSports() {
	f := new(adidas)
	f.makeShirts()
	f.makeShoes()
}
