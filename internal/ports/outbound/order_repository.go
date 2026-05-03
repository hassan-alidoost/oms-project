package outbound

type OrderRepository interface {
    Save() error
	Find() error
}