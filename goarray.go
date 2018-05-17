package main
import "fmt"
func main(){
  var twoD [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i - j
        }
    }
    fmt.Println("2d: ", twoD)
}
