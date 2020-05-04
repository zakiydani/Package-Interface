package sepeda

type Maju interface {
	Cepat() float32
	Lambat() float32
}
type Sepeda struct {
	jumlahBan, jumlahGigi int
	warna string
}

func CreateSepeda(ban, gigi int, warna string) *Sepeda {
	return &Sepeda{
		jumlahBan:	ban,
		jumlahGigi: gigi,
		warna:		warna,
		
	}
}

var waktuSepeda float32

func (m *Sepeda) WaktuTempuh(jarak float32) float32 {
	waktuSepeda := jarak * 2.5
	return waktuSepeda
}

func (m *Sepeda) Cepat(waktuSepeda float32) float32 {
	return waktuSepeda * 2
}

func (m *Sepeda) Lambat(waktuSepeda float32) float32 {
	return waktuSepeda * 0.5
}