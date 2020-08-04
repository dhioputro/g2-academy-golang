package main

import (
	"fmt"
	"math"
)

type Basic struct {
	n1, n2 int
}
type Single struct {
	n1 int
}
type FloatSingle struct {
	n1 float64
}
type FloatTab struct {
	n1, n2 float64
}
type Volume struct {
	n1, n2, n3 int
}
type FloatVol struct {
	n1, n2, n3 float64
}

func main() {
	var input string
	for input != "99" {
		fmt.Println("What to do?")
		fmt.Println("1. Additional")
		fmt.Println("2. Minus")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Root")
		fmt.Println("6. Squared")
		fmt.Println("7. Luas Persegi")
		fmt.Println("8. Luas Lingkaran")
		fmt.Println("9. Volume Tabung")
		fmt.Println("10. Volume Balok")
		fmt.Println("11. Volume Prisma")
		fmt.Println("99. end")
		fmt.Println("=============")
		// scanner := bufio.NewScanner(input)
		// input, _ := scanner.ReadString('\n')
		fmt.Scanln(&input)

		switch input {
		case "1":
			var b1 int
			var b2 int
			fmt.Println("Input first Number:")
			fmt.Scanln(&b1)
			fmt.Println("Input second Number:")
			fmt.Scanln(&b2)
			kalkulasi := Basic{b1, b2}
			fmt.Println("Hasil Tambah:", kalkulasi.Additional())

		case "2":
			var b1 int
			var b2 int
			fmt.Println("Input first Number:")
			fmt.Scanln(&b1)
			fmt.Println("Input second Number:")
			fmt.Scanln(&b2)
			kalkulasi := Basic{b1, b2}
			fmt.Println("Hasil Kurang:", kalkulasi.Minus())

		case "3":
			var b1 int
			var b2 int
			fmt.Println("Input first Number:")
			fmt.Scanln(&b1)
			fmt.Println("Input second Number:")
			fmt.Scanln(&b2)
			kalkulasi := Basic{b1, b2}
			fmt.Println("Hasil Kali:", kalkulasi.Multiply())

		case "4":
			var b1 int
			var b2 int
			fmt.Println("Input first Number:")
			fmt.Scanln(&b1)
			fmt.Println("Input second Number:")
			fmt.Scanln(&b2)
			kalkulasi := Basic{b1, b2}
			fmt.Println("Hasil Bagi:", kalkulasi.Divide())

		case "5":
			fmt.Println("Input Number:")
			var r float64
			fmt.Scanln(&r)
			kalkulasi := FloatSingle{r}
			fmt.Println("Hasil Bagi:", kalkulasi.Root())

		case "6":
			fmt.Println("Input Number:")
			var b1 int
			fmt.Scanln(&b1)
			kalkulasi := Single{b1}
			fmt.Println("Hasil Kuadrat:", kalkulasi.Squared())

		case "7":
			var p int
			var l int
			fmt.Println("Input panjang:")
			fmt.Scanln(&p)
			fmt.Println("Input lebar:")
			fmt.Scanln(&l)
			kalkulasi := Basic{p, l}
			fmt.Println("Luas Persegi:", kalkulasi.Persegi())

		case "8":
			fmt.Println("Input Jari -jari:")
			var r float64
			fmt.Scanln(&r)
			kalkulasi := FloatSingle{r}
			fmt.Println("Luas Lingkaran:", kalkulasi.Lingkaran())

		case "9":
			var r float64
			var t float64
			fmt.Println("Input jari-jari:")
			fmt.Scanln(&r)
			fmt.Println("Input tinggi:")
			fmt.Scanln(&t)
			kalkulasi := FloatTab{r, t}
			fmt.Println("Volume Tabung:", kalkulasi.Tabung())

		case "10":
			var p int
			var l int
			var t int
			fmt.Println("Input panjang:")
			fmt.Scanln(&p)
			fmt.Println("Input lebar:")
			fmt.Scanln(&l)
			fmt.Println("Input tinggi:")
			fmt.Scanln(&t)
			kalkulasi := Volume{p, l, t}
			fmt.Println("Volume Balok:", kalkulasi.Balok())
		case "11":
			var p float64
			var l float64
			var t float64
			fmt.Println("Input panjang:")
			fmt.Scanln(&p)
			fmt.Println("Input lebar:")
			fmt.Scanln(&l)
			fmt.Println("Input tinggi:")
			fmt.Scanln(&t)
			kalkulasi := FloatVol{p, l, t}
			fmt.Println("Volume Prisma:", kalkulasi.Prism())

		case "99":
			fmt.Println("Please input correct number")
		}
	}

}

func (c Basic) Additional() int {
	return c.n1 + c.n2
}
func (c Basic) Minus() int {
	return c.n1 - c.n2
}
func (c Basic) Multiply() int {
	return c.n1 * c.n2
}
func (c Basic) Divide() int {
	return c.n1 / c.n2
}
func (i FloatSingle) Root() float64 {
	return math.Sqrt(i.n1)
}
func (i Single) Squared() int {
	return i.n1 * i.n1
}
func (c Basic) Persegi() int {
	return c.n1 * c.n2
}
func (i FloatSingle) Lingkaran() float64 {
	return i.n1 * i.n1 * 3.14
}
func (f FloatTab) Tabung() float64 {
	return f.n1 * f.n1 * 3.14 * f.n2
}
func (v Volume) Balok() int {
	return v.n1 * v.n2 * v.n3
}
func (v FloatVol) Prism() float64 {
	return v.n1 * v.n2 * v.n3 * 0.5
}
