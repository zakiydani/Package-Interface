package motor

type Maju interface {
	cepat() float32
	lambat() float32
}

type Motor struct {
	jumlahBan, jumlahGigi int
	warna string
}

func CreateMotor(ban, gigi int, warna string) *Motor {
	return &Motor{
		jumlahBan:	ban,
		jumlahGigi: gigi,
		warna:		warna,
	}

}

var waktuMotor float32

func (m *Motor) waktuTempuh(jarak float32) float32 {
	waktuMotor := jarak * 0.5
	return waktuMotor
}

func (m *Motor) Cepat() float32 {
	return waktuMotor * 0.5
}

func (m *Motor) Lambat(jarak float32) float32 {
	return waktuMotor * 2
}