package main

import "fmt"

func main() {
	printNominal()
	printTypeData()
	printPercent()
	nilaiBoolean()
	unicodeRusia()
	nilaiBase1021()
	nilaiBase825()
	nilaiBase16Kecil()
	nilaiBase16Besar()
	dalamBahasaRusia1()
	nilaiFloat64()
	NilaiFloatScientific()
}

func printNominal() {
	var nominal = 21
	fmt.Printf("%d\n", nominal) // %d akan mencetak angka dalam format desimal
}

func printTypeData() {
	var i = 42 // %T akan mencetak tipe data dari i
	fmt.Printf("%T\n", i)
}

func printPercent() {
	fmt.Printf("%%\n") // %% akan mencetak tanda persen (%)
}
func nilaiBoolean() {
	var isTrue = true
	fmt.Printf("%t\n\n", isTrue) // %t akan mencetak nilai boolean
}
func unicodeRusia() {
	var rusia = 'я'
	fmt.Printf("%c\n", rusia) // %c akan mencetak karakter Unicode
}
func nilaiBase1021() {
	angka := 21
	// Menampilkan dalam base 10 (desimal)
	fmt.Printf("%d\n", angka)
}
func nilaiBase825() {
	angka := 25
	// Menampilkan dalam base 10 (desimal)
	fmt.Printf("%d\n", angka)
}
func nilaiBase16Kecil() {
	fmt.Printf("%x\n", 15)
}
func nilaiBase16Besar() {
	fmt.Printf("%X\n", 15)
}

func dalamBahasaRusia1() {
	fmt.Printf("%U\n\n", 'Я')
}
func nilaiFloat64() {
	var k float64 = 123.456
	fmt.Printf("%f\n", k)
}

func NilaiFloatScientific() {
	var k float64 = 123.456
	fmt.Printf("%E\n", k)
}
