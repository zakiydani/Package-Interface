package motor

// import (
// 	"fmt"
// 	"github.com/zakiydani/project/contract"
// ) 

type Maju interface {
	Cepat() float32
	Lambat() float32
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

func (m *Motor) WaktuTempuh(jarak float32) float32 {
	waktuMotor := jarak * 5
	return waktuMotor
}

func (m *Motor) Cepat(waktuMotor float32) float32 {
	return waktuMotor * 2
}

func (m *Motor) Lambat(waktuMotor float32) float32 {
	return waktuMotor * 0.5
}

// func Hitung(h contract.Maju) {
// 	fmt.Println(h.Cepat())
// }

// func Motors() {
// 	Hitung()
// }