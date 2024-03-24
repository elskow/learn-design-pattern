package AbstractFactory

type kondisi string

const (
	Dingin kondisi = "Dingin"
	Panas  kondisi = "Panas"
)

type IMinuman interface {
	setNama(nama string)
	setPrice(price int64)
	setKondisi(kondisi kondisi)
	getNama() string
	getPrice() int64
	getKondisi() kondisi
}

type Minuman struct {
	nama    string
	price   int64
	kondisi kondisi
}

func (m *Minuman) setNama(nama string) {
	m.nama = nama
}

func (m *Minuman) setPrice(price int64) {
	m.price = price
}

func (m *Minuman) setKondisi(kondisi kondisi) {
	m.kondisi = kondisi
}

func (m *Minuman) getNama() string {
	return m.nama
}

func (m *Minuman) getPrice() int64 {
	return m.price
}

func (m *Minuman) getKondisi() kondisi {
	return m.kondisi
}
