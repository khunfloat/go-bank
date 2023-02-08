package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "71110", DateOfBirth: "1978-12-15", Status: 1},
		{CustomerID: 1002, Name: "Chayoot", City: "Bangkok", ZipCode: "10600", DateOfBirth: "2003-12-15", Status: 0},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
