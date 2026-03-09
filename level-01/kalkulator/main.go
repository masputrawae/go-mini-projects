package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrZeroDivision = errors.New("Tidak bisa di bagi dengan 0")

func main() {
	// Input angka
	a := input("Masukkan angka pertama: ")
	b := input("Masukkan angka pertama: ")

	// Tampilkan opsi
	fmt.Println("====================================")
	fmt.Println("[1] Penjumlahan")
	fmt.Println("[2] Pengurangan")
	fmt.Println("[3] Perkalian")
	fmt.Println("[4] Pembagian")
	fmt.Println("[5] Modulus (sisa bagi)")
	fmt.Println("====================================")

	// Ambil input opsi
	opt := input("Pilih Operasi (1, 2, 3, 4, 5): ")

	// Operasi
	switch opt {
	case 1:
		fmt.Printf("Hasil dari (%v + %v) adalah: %v\n", a, b, a + b)
	case 2:
		fmt.Printf("Hasil dari (%v - %v) adalah: %v\n", a, b, a - b)
	case 3:
		fmt.Printf("Hasil dari (%v * %v) adalah: %v\n", a, b, a * b)
	case 4:
		// Cek zero division
		if b == 0 {
			fmt.Println(ErrZeroDivision)
			return
		}
		fmt.Printf("Hasil dari (%v / %v) adalah: %v\n", a, b, a / b)
	case 5:
		// Cek zero division
		if b == 0 {
			fmt.Println(ErrZeroDivision)
			return
		}
		fmt.Printf("Hasil dari (%v %% %v) adalah: %v\n", a, b, math.Mod(a, b))
	}
}

// Fungsi untuk menerima input
func input(message string) float64 {
	var num float64
	fmt.Printf("%s", message)
	fmt.Scan(&num)
	return num
}
