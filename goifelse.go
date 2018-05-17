package main
import "fmt"
func main(){

fmt.Println("enter the number to be checked")
var n int
fmt.Scanln(&n)
if n%2 == 0 {
  fmt.Println("even")
} else {
  fmt.Println("odd")
}
}
