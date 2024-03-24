package AbstractFactory

type IMakanan interface {
	setNama(nama string)
	setPrice(price int64)
	getNama() string
	getPrice() int64
}

type Makanan struct {
	nama  string
	price int64
}

func (m *Makanan) setNama(nama string) {
	m.nama = nama
}

func (m *Makanan) setPrice(price int64) {
	m.price = price
}

func (m *Makanan) getNama() string {
	return m.nama
}

func (m *Makanan) getPrice() int64 {
	return m.price
}
