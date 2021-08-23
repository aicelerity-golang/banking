package domain

type CustomerRepositoryStub struct {
	customers []Customers
}

func (s CustomerRepositoryStub) FindAll() ([]Customers, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customers{
		{10001, "Raj", "Bangalore", "56001", "1995-04-02", 1},
		{10002, "Kavi", "Bangalore", "56002", "1990-06-01", 1},
		{10003, "Som", "Bangalore", "56003", "1996-08-04", 1},
		{10004, "Vittal", "Bangalore", "56001", "1993-05-12", 1},
	}
	return CustomerRepositoryStub{customers}
}
