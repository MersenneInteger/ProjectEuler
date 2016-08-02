package main
import "fmt"
func main(){

    var target, count int64 = 0,2
    
    for count = 2; target < 10002; count++ {

        if (isPrime(count)) == true {
            target++      
            if target == 10001 {
                fmt.Print(count)
                break
            }    
        }
    }
}
func isPrime(n int64) bool {

    var i int64 = 2; 
    for i = 2; i <= (n/2); i++ {
        if n % i == 0 { return false }
    }
    return true
}
