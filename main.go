// package main

// import (
// 	"fmt"
// 	// "github.com/zakiydani/project"
// )

// type Sepeda struct {
// 	jumlahBan, jumlahGigi int
// 	warna string
// }

// type Motor struct {
// 	jumlahBan, jumlahGigi int
// 	warna string
// }

// func (s *Sepeda) waktuTempuh(jarak float32) float32 {
// 	waktu := jarak * 2.5 
// 	return waktu
// }

// func (m *Motor) waktuTempuh(jarak float32) float32 {
// 	waktu := jarak * 5 
// 	return waktu
// }

// type Maju interface {
// 	cepat() float32
// 	lambat() float32
// }

// func (m Motor) cepat() float32 {
// 	return  m * 2
// }

// func (m Motor) lambat() float32 {
// 	return m / 2
// }

// func main() {
// 	var kecepatan Maju

// 	kecepatan = Motor{20}
// }