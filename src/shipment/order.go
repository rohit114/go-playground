package shipment

type Order struct {
	id    int
	item  Product
	state State
}

type OrderObject struct {
	Id    int
	Item  any
	State State
}

// Getter methods to access private fields
func (o *Order) GetID() int {
	return o.id
}

func (o *Order) GetItem() Product {
	return o.item
}

func (o *Order) GetState() State {
	return o.state
}

func CreateNewOrder(id int, item Product) (*Order, OrderObject) {
	order := Order{id: id, item: item, state: ORDER_PLACED}

	return &order, GetOrderObject(&order)
}

func GetOrderObject(order *Order) OrderObject {

	return OrderObject{Id: order.id, Item: ProductObj{Id: order.item.id, Name: order.item.name, Price: order.item.price}, State: order.state}
}

func (o *Order) UpdateStateOfOrder(new_state State) {
	o.state = new_state
}
