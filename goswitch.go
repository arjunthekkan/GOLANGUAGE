package main
import "fmt"
func main(){
  fmt.Println("enter the two numbers")
  var a,b float64
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  fmt.Println("operation to be execute\naddition: 1 \n subtraction: 2 \n multiplication: 3 \n division: 4")
  var i int
  fmt.Scanln(&i)
  switch i {
  case 1:
    fmt.Println(a+b)
  case 2:
    fmt.Println(a-b)
  case 3:
    fmt.Println(a*b)
  case 4:
    fmt.Println(float64(a)/b)
    }
}
