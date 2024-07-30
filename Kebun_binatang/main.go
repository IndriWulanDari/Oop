// package main

// import "fmt"

// // Class Objek Animal
// type Animal struct (){
// 	Nama string
// 	Spesies string
// 	Usia int
// }

// func (Animal) Suara() string {
// 	return "Suaranya"
// }

// //Membuat Struk/Object Hewan hewan
// //*Method Suara
// //Simba the Lion says Roar!
// //Dumbo the Elephant says Trumpet!
// //George the Monkey says Ooh Ooh Aah Aah!

// type Singa struct {
// 	Animal
// 	Color string
// }

// func (Singa1 Singa) Suara() string{
// 	return "Roar!"
// }

// type Gajah struct {
// 	Animal
// 	Color string
// }

// func (Gajah1 Gajah Suara) string{
// 	return "Trumpet! "
// }

// type Monyet struct {
// 	Animal
// 	Color string
// }

// func (Monyet1 Monyet) Suara() string{
// 	return "Ooh Ooh Aah Aah!"
// }

// func main() {
// 	Singa1 := Singa{
// 		Animal : Animal{Nama: "Lion", Spesies: "Mamalia", Umur: "2"},
// 		Color: "Orange"
// 	}

// 	Gajah1 := Gajah{
// 		Animal : Animal{Nama: "Lion", Spesies: "Mamalia", Umur: "2"},
// 		Color: "Orange"
// 	}

// 	Monyet1 := Monyet{
// 		Animal : Animal{Nama: "Lion", Spesies: "Mamalia", Umur: "2"},
// 		Color: "Orange"
// 	}

// 	fmt.Println(Singa1.Nama, "suara", Singa1.Suara, "dan warnanya", Singa1.Color)
// 	fmt.Println(Gajah1.Nama, "suara", Gajah1.Suara, "dan warnanya", Gajah1.Color)
// 	fmt.Println(Monyet1.Nama, "suara", Monyet1.Suara, "dan warnanya", Monyet1.Color)
// }

package main

import "fmt"

// Class Objek Animal
type Animal struct {
	Nama    string
	Spesies string
	Usia    int
}

func (Animal) Suara() string {
	return "Suaranya"
}

// Membuat Struk/Object Hewan hewan
// Method Suara
// Simba the Lion says Roar!
// Dumbo the Elephant says Trumpet!
// George the Monkey says Ooh Ooh Aah Aah!

type Singa struct {
	Animal
	Color string
}

func (Singa) Suara() string {
	return "Roar!"
}

type Gajah struct {
	Animal
	Color string
}

func (Gajah) Suara() string {
	return "Trumpet!"
}

type Monyet struct {
	Animal
	Color string
}

func (Monyet) Suara() string {
	return "Ooh Ooh Aah Aah!"
}

func main() {
	Singa1 := Singa{
		Animal: Animal{Nama: "Simba", Spesies: "Mamalia", Usia: 2},
		Color:  "Orange",
	}

	Gajah1 := Gajah{
		Animal: Animal{Nama: "Dumbo", Spesies: "Mamalia", Usia: 10},
		Color:  "Gray",
	}

	Monyet1 := Monyet{
		Animal: Animal{Nama: "George", Spesies: "Mamalia", Usia: 5},
		Color:  "Brown",
	}

	fmt.Println(Singa1.Nama, "suara", Singa1.Suara(), "dan warnanya", Singa1.Color)
	fmt.Println(Gajah1.Nama, "suara", Gajah1.Suara(), "dan warnanya", Gajah1.Color)
	fmt.Println(Monyet1.Nama, "suara", Monyet1.Suara(), "dan warnanya", Monyet1.Color)
}
