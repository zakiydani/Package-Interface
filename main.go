package main

import (
	"fmt"
	"github.com/zakiydani/project/motor"
	"github.com/zakiydani/project/sepeda"
)

func main() {
	// ArrMotor := new(motor.Motor)
	fmt.Println("=================Motor==================")
	m := motor.CreateMotor(2, 4, "Hitam")
	jarakMotor := m.WaktuTempuh(50)
	fmt.Println("Motor dengan data:",m)
	fmt.Println("Memiliki kecepatan Awal:", jarakMotor)
	fmt.Println("Kemudian mempercepat lajunya menjadi:", m.Cepat(jarakMotor))
	fmt.Println("dan diperlambat lajunya menjadi:", m.Lambat(jarakMotor))
	fmt.Println("========================================")

	fmt.Println("\n=================Sepeda=================")
	s := sepeda.CreateSepeda(2, 2, "Putih")
	jarakSepeda := s.WaktuTempuh(50)
	fmt.Println("Sepeda dengan data:", s)
	fmt.Println("Memiliki kecepatan Awal:", jarakSepeda)
	fmt.Println("Kemudian mempercepat lajunya menjadi:", s.Cepat(jarakSepeda))
	fmt.Println("dan diperlambat lajunya menjadi:", s.Lambat(jarakSepeda))
	fmt.Println("========================================")

}
