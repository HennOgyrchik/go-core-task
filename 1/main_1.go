package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {

	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Print("2. ")
	for _, a := range []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum} {
		printType(a)
	}

	s := numbers{
		d: numDecimal,
		o: numOctal,
		h: numHexadecimal,
		f: pi,
		s: name,
		b: isActive,
		c: complexNum,
	}.String()

	fmt.Printf("3. %s\n", s)

	fmt.Printf("4. %v\n", toSliceRune(s))

	fmt.Print("5. ")
	printHash(toSliceRune(s), "go-2024")
}

func printType(a any) {
	fmt.Printf("This is %T\n", a)
}

type numbers struct {
	d int
	o int
	h int
	f float64
	s string
	b bool
	c complex64
}

func (n numbers) String() string {
	return fmt.Sprintf("%d, 0%o, 0x%X, %.2f, %s, %t, %v", n.d, n.o, n.h, n.f, n.s, n.b, n.c)
}

func toSliceRune(s string) []rune {
	return []rune(s)
}

func printHash(r []rune, salt string) {
	hasher := sha256.New()

	m := len(r) / 2

	hasher.Write([]byte(string(append(r[:m], append([]rune(salt), r[m:]...)...))))

	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}
