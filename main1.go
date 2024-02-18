package main

import (
	"fmt"
)

func yayan() {
	fmt.Println("Hello Wohrld!") //print
	a := 5
	b := 7
	c := sum(a, b)
	fmt.Printf("Jumlahnya adalah %d%s\n", c, "yayan")
	mult := func(x, y, z int) (hasil int) {
		hasil = x * y * z
		return
	}
	x := 5
	y := 6
	z := 7
	hasil := mult(x, y, z)
	fmt.Printf("Hasil kalinya adalah %d\n", hasil)
	palindrom("ziadai")
}

func sum(h, k int) (jml int) {
	jml = h + k
	return
}

func palindrom(m string) {
	// for _, char := range m{

	// }
	z := true
	for i := 0; i < len(m); i++ {
		if m[i] != m[len(m)-i-1] {
			z = false
		}
	}
	if z {
		fmt.Println("Ini Palindrom")
	} else {
		fmt.Println("Ini bukan Palindrom")
	}
}
