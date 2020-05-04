package main

import (
	"fmt"
	"github.com/zakiydani/project/motor"
)

func main() {
	ArrMotor := new(motor.Motor)
	ArrMotor1 := motor.CreateMotor(2, 3, "Hitam")
	jarakMotor := ArrMotor.waktuTempuh(20)
	fmt.Println(ArrMotor)
	// fmt.Println(ArrMotor.cepat(jarakMotor+))
}
