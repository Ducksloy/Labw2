package main
import "fmt"

func main() {
  var a int
  var b int
  // takes input value for name
  fmt.Print("Enter first number: ")
  fmt.Scan(&a)
  fmt.Print("Enter second number: ")
  fmt.Scan(&b)

  if a>0 && b<0{
  fmt.Printf("%d is positive\n",a)
  fmt.Printf("%d is Negative",b)
  }
  if  a<0 && b>0{
  fmt.Printf("%d is positive\n",b)
  fmt.Printf("%d is Negative",a)
  }
  if a>0 && b>0{
  fmt.Printf("%d and %d are possitive",a,b)
  }
  if a>b || b>a{
  fmt.Printf("\nleast one number is greater:%d,%d\n",a,b)
  }
  if !(a== b){
		fmt.Printf("%d and %d are not equal", a,b)
	}
	if (a== b){
		fmt.Printf("\n%d and%d are equal", a,b)
	}
}