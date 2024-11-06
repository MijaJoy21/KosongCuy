package gudang

type Gudang struct {
	Name   string
	Barang string
	Harga  int
	Stok   int
}

type Gudangs interface {
	ShowStok() int
	ShowName() string
	ShowBarang() string
	ShowHarga() int
}

func InitGudang(name string, barang string, harga int, stok int) Gudangs {
	return &Gudang{
		Name:   name,
		Barang: barang,
		Harga:  harga,
		Stok:   stok,
	}
}

func (g *Gudang) ShowStok() int {
	return g.Stok
}

func (g *Gudang) ShowName() string {
	return g.Name
}

func (g *Gudang) ShowBarang() string {
	return g.Barang
}

func (g *Gudang) ShowHarga() int {
	return g.Harga
}
