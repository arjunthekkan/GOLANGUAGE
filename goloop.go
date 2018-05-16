package main
import "fmt"

func main(){
    var n,sum int
    fmt.Println("Enter a positive integer : ")
    fmt.Scanln(&n)
    for i:=1; i<=n; i++ {
        sum = sum+i
    }
    fmt.Println("Sum : ",sum)
}
