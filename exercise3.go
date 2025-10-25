package main

import "fmt"

func myXOR(a, b int) {
	result := a ^ b
	fmt.Printf("%d XOR %d = %d\n", a, b, result)
}

func myNOT(a int) {
	result := ^a
	fmt.Printf("NOT %d = %d\n", a, result)
}

func myOR(a, b int) {
	result := a | b
	fmt.Printf("%d OR %d = %d\n", a, b, result)
}

func myAND(a, b int) {
	result := a & b
	fmt.Printf("%d AND %d = %d\n", a, b, result)
}

func myLeftShift(a int, shift uint) {
	result := a << shift
	fmt.Printf("%d << %d = %d\n", a, shift, result)
}

func myRightShift(a int, shift uint) {
	result := a >> shift
	fmt.Printf("%d >> %d = %d\n", a, shift, result)
}

func main() {
	a := 12
	b := 5

	fmt.Println("Bitwise Operations:")
	myXOR(a, b)
	myNOT(a)
	myOR(a, b)
	myAND(a, b)
	myLeftShift(a, 2)
	myRightShift(a, 2)

	fmt.Println("\nAssignment Operators:")

	c := 10
	fmt.Printf("Initial c = %d\n", c)

	c &= 3
	fmt.Printf("c &= 3 -> %d\n", c)

	c |= 6
	fmt.Printf("c |= 6 -> %d\n", c)

	c ^= 2
	fmt.Printf("c ^= 2 -> %d\n", c)

	c <<= 1
	fmt.Printf("c <<= 1 -> %d\n", c)

	c >>= 2
	fmt.Printf("c >>= 2 -> %d\n", c)
}
