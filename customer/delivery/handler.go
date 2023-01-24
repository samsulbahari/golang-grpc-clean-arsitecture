package delivery

import (
	"context"
	"grpc/customer/delivery/customer_grpc"
	"grpc/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CustomerHandler struct {
	customerServ domain.CustomerService
}

func NewCustomerHandler(cs domain.CustomerService, gserver *grpc.Server) {
	custumerServ := &CustomerHandler{
		customerServ: cs,
	}
	customer_grpc.RegisterCustomerHandlerServer(gserver, custumerServ)
	reflection.Register(gserver)

}
func (ch *CustomerHandler) transformCustomerRPC(ar *domain.Customer) *customer_grpc.Customer {

	if ar == nil {
		return nil
	}

	res := &customer_grpc.Customer{
		Id:      ar.Id,
		Name:    ar.Name,
		Address: ar.Address,
		Email:   ar.Email,
		Phone:   ar.Phone,
	}
	return res

}

func (ch *CustomerHandler) transformArticleData(ar *customer_grpc.Customer) *domain.Customer {
	res := &domain.Customer{
		Id:      ar.Id,
		Name:    ar.Name,
		Address: ar.Address,
		Email:   ar.Email,
		Phone:   ar.Phone,
	}
	return res
}

func (ch *CustomerHandler) GetCustomer(ctx context.Context, ce *customer_grpc.Empty) (*customer_grpc.Customers, error) {
	// data := ctx.Value("tokenInfo")
	// fmt.Println(data)
	customer, err := ch.customerServ.GetData()
	if err != nil {
		return nil, err
	}
	arrCustomer := make([]*customer_grpc.Customer, len(customer))

	for i, a := range customer {
		ar := ch.transformCustomerRPC(&a)
		arrCustomer[i] = ar
	}
	result := &customer_grpc.Customers{
		Message: "succes get data",
		Data:    arrCustomer,
	}
	return result, nil
}
func (ch *CustomerHandler) InsertCustomer(ctx context.Context, cus *customer_grpc.Customer) (*customer_grpc.CustomerResponse, error) {
	cs := ch.transformArticleData(cus)
	id, err := ch.customerServ.Insert(*cs)
	if err != nil {
		return nil, err
	}

	return &customer_grpc.CustomerResponse{
		Id: id,
	}, nil
}
