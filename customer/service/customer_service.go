package service

import (
	"grpc/domain"
)

type CustomerService struct {
	customerRepo domain.CustomerRepo
}

func NewCustomerService(cs domain.CustomerRepo) *CustomerService {
	return &CustomerService{customerRepo: cs}

}

func (cs *CustomerService) GetData() ([]domain.Customer, error) {
	customer, err := cs.customerRepo.GetData()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cs *CustomerService) Insert(customer domain.Customer) (int64, error) {
	id, err := cs.customerRepo.Insert(customer)
	if err != nil {
		return 0, err
	}
	return id, nil
}
