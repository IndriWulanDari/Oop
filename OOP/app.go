// Contoh Pengaplikasian OOP
package main

import "fmt"

// Definisi kelas atau struct Hewan
type Animal struct {
	Name   string
	Age    int
	Weight float64
}

// Metode untuk kelas Hewan
func (a Animal) Eat() {
	fmt.Printf("%s is eating.\n", a.Name)
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping.\n", a.Name)
}

// Definisi kelas Singa yang mewarisi dari kelas Hewan
type Lion struct {
	Animal // Pewarisan
}

// Metode khusus untuk kelas Singa
func (l Lion) Roar() {
	fmt.Printf("%s is roaring.\n", l.Name)
}

func main() {
	// Membuat objek Singa dari kelas Hewan
	simba := Lion{Animal{Name: "Simba", Age: 5, Weight: 200}}
	simba.Eat()   // Memanggil metode yang diwarisi dari Hewan
	simba.Sleep() // Memanggil metode yang diwarisi dari Hewan
	simba.Roar()  // Memanggil metode khusus dari Singa
}
