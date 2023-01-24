package domain

type Customer struct {
	Id      int64
	Name    string
	Address string
	Email   string
	Phone   string
}

type CustomerRepo interface {
	GetData() ([]Customer, error)
	Insert(customer Customer) (int64, error)
}

type CustomerService interface {
	GetData() ([]Customer, error)
	Insert(customer Customer) (int64, error)
}
