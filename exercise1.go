package main
import "fmt"

func main() {
  var a int
  var b int
  
  fmt.Print("Enter first number: ")
  fmt.Scan(&a)
  fmt.Print("Enter second number: ")
  fmt.Scan(&b)
  
a = b
  fmt.Printf("After Assignment the first number is: %d\n", a)
a += b
  fmt.Printf("After addition Assignment the first number is: %d\n", a)
a -= b
  fmt.Printf("After subtraction Assignment the first number is: %d\n", a)
a *= b
  fmt.Printf("After Multiple Assignment the first number is: %d\n", a)
a /= b
  fmt.Printf("After Devide Assignment the first number is: %d\n", a)
a %= b
  fmt.Printf("After Modulus Assignment the first number is: %d\n", a)

}