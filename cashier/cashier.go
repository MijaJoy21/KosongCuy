package cashier

import (
	"fmt"
	cs "tugas/customer"
	g "tugas/gudang"
)

type Cashier struct {
	Name string
	CS   cs.Customers
	GD   g.Gudangs
}

type Cashiers interface {
	Pembelian()
	CheckNamaKasir() string
}

func InitCashier(name string, newCs cs.Customers, newG g.Gudangs) Cashiers {
	return &Cashier{
		Name: name,
		CS:   newCs,
		GD:   newG,
	}
}

func (c *Cashier) Pembelian() {
	stock := c.GD.ShowStok()
	quantity := 1

	if stock < 1 {
		fmt.Println("Stock Habis")
	}

	if stock < quantity {
		fmt.Println("Stock Kurang")
	}

	totalHarga := c.GD.ShowHarga() * quantity

	if totalHarga > c.CS.CheckSaldo() {
		fmt.Println("Saldo Tidak Cukup")
	}

	sisaSaldo := c.CS.CheckSaldo() - totalHarga

	fmt.Println("Dari Gudang : ", c.GD.ShowName())
	fmt.Println("Nama Product: ", c.GD.ShowBarang())
	fmt.Println("Harga Product: ", c.GD.ShowHarga())
	fmt.Println("Quantity : ", quantity)
	fmt.Println("Total Pembelian : ", totalHarga)
	fmt.Println("")
	fmt.Println("Nama Pembeli : ", c.CS.CheckName())
	fmt.Println("Sisa Saldo : ", sisaSaldo)

}

func (c *Cashier) CheckNamaKasir() string {
	return c.Name
}
