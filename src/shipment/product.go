package shipment

type Product struct {
	id    int
	name  string
	price int
}

type ProductObj struct {
	Id    int
	Name  string
	Price int
}

func (p *Product) Update(id int, name string, price int) {
	p.id = id
	p.name = name
	p.price = price
}

func CreateNewProduct(id int, name string, price int) *Product {
	return &Product{id: id, name: name, price: price}
}
