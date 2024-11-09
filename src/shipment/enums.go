package shipment

type State string
type Event string

const (
	ORDER_PLACED     State = "ORDER_PLACED"
	ORDER_PACKED     State = "ORDER_PACKED"
	ORDER_PICKED_UP  State = "ORDER_PICKED_UP"
	ORDER_SHIPPED    State = "ORDER_SHIPPED"
	ORDER_IN_TRANSIT State = "ORDER_IN_TRANSIT"
	ORDER_DELIVERED  State = "ORDER_DELIVERED"
)

const (
	PACKED     Event = "PACKED"
	PICKED_UP  Event = "PICKED_UP"
	SHIPPED    Event = "SHIPPED"
	IN_TRANSIT Event = "IN_TRANSIT"
	DELIVERED  Event = "DELIVERED"
)
