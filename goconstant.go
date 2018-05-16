package main
import "fmt"
import "math"
const s string = "CONSTANT"
func main(){
  fmt.Println(s)
  const n = 100000
  const d = 13e10 / n
  fmt.Println(d)
  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))


}
