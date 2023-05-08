package abstract_factory

type adidas struct {
}

type adidasShoe struct {
	brand string
	shoe
}

type adidasShirt struct {
	brand string
	shirt
}

func (a *adidas) makeShoes() iShoe {
	return &adidasShoe{
		brand: "adidas",
		shoe: shoe{
			color: "red",
			size:  14,
			price: 100,
		},
	}
}

func (a *adidas) makeShirts() iShirt {
	return &adidasShirt{
		brand: "adidas",
		shirt: shirt{
			size:     10,
			color:    "blue",
			price:    119,
			material: "s",
		},
	}
}
