package main

import "fmt"

func islemHavuz(birinciSayi int, ikinciSayi int) (int, int, int, int) {
	toplamaSonuc := birinciSayi + ikinciSayi
	cikarmaSonuc := birinciSayi - ikinciSayi
	carpmaSonuc := birinciSayi * ikinciSayi
	bolmeSonuc := birinciSayi / ikinciSayi
	return toplamaSonuc, cikarmaSonuc, carpmaSonuc, bolmeSonuc
}

func main() {
	fmt.Println("Go hesap makinesine hoş geldiniz. Hesap makinesini kullanabilmek için iki adet sayı giriniz.")
	fmt.Println("Lütfen birinci sayıyı giriniz.")

	var birinciSayi int
	fmt.Scan(&birinciSayi)
	fmt.Println("Lütfen ikinci sayıyı giriniz.")
	var ikinciSayi int
	fmt.Scan(&ikinciSayi)

	_toplamaSonuc, _cikarmaSonuc, _carpmaSonuc, _bolmeSonuc := islemHavuz(birinciSayi, ikinciSayi)
	fmt.Println("Toplama işleminin sonucu", _toplamaSonuc)
	fmt.Println("Çıkartma işleminin sonucu", _cikarmaSonuc)
	fmt.Println("Çarpma işleminin sonucu", _carpmaSonuc)
	fmt.Println("Bölme işleminin sonucu", _bolmeSonuc)
}

//AEC 09.21.17
//Türkiye Havasahasında bir yerde
