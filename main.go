/* Contoh OOP
class Mobil {
	String mark;
	int tahun;

	Mobil(this.merk, this.tahun);

	void infoMobil(){
		print("Mobil $merk tahun $tahun");
	}
}

void main(){
	Mobil mobil1 = Mobil("Toyota", 2020);
	Mobil mobil2 = Mobil("Honda", 2018);

	mobil1.infoMobil();
	mobil2.infoMobil();

}
*/

package main

import "fmt"

// Mendefinisikan struct Mobil
type Mobil struct {
	merk  string
	tahun int
}

// Membuat constructor untuk struct Mobil
func NewMobil(merk string, tahun int) Mobil {
	return Mobil{merk: merk, tahun: tahun}
}

// Mendefinisikan method untuk struct Mobil
func (m Mobil) infoMobil() {
	fmt.Printf("Mobil %s tahun %d\n", m.merk, m.tahun)
}

func main() {
	// Membuat instance dari struct Mobil
	mobil1 := NewMobil("Toyota", 2020)
	mobil2 := NewMobil("Honda", 2018)

	// Memanggil method infoMobil untuk masing-masing instance
	mobil1.infoMobil()
	mobil2.infoMobil()
}


