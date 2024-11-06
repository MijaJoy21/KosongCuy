package main

import (
	cr "tugas/cashier"
	cs "tugas/customer"
	g "tugas/gudang"
)

func main() {
	newCustomer := cs.InitCustomer("Mija", 300000)
	newGudang := g.InitGudang("KINS", "Smartwatch", 75000, 3)
	newCashier := cr.InitCashier("Lara", newCustomer, newGudang)

	newCashier.Pembelian()

}
