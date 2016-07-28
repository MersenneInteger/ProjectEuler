package main
import "fmt"
func main(){

    var sumOfSquare, squareOfSum, i uint32 = 0,0,0

    for i = 1; i < 101; i++ {
       sumOfSquare += i * i
       squareOfSum += i
    }
    squareOfSum *= squareOfSum
    fmt.Print(squareOfSum - sumOfSquare) 
}
