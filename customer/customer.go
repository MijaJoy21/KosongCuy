package swalayan

type Customer struct {
	Name  string
	Saldo int
}

type Customers interface {
	CheckSaldo() int
	CheckName() string
}

func InitCustomer(name string, saldo int) Customers {
	return &Customer{
		Name:  name,
		Saldo: saldo,
	}
}

func (c *Customer) CheckSaldo() int {
	return c.Saldo
}

func (c *Customer) CheckName() string {
	return c.Name
}
