package main
import "fmt"
func main(){
    
    var count, limit, i, sum int64 = 2, 2000000, 0, 0
    var list [2000000]int64

    for count = 2; count < limit; count++ {
        list[count] = 1
    }

    for count = 2; count < limit; count++{
        for i = count; i * count < limit; i++{
            list[ i * count] = 0        
        }        
    }
    for count = 2; count < limit; count++{
         if list[count] == 1 {
            sum += count        
        }        
    }

    fmt.Print(sum)
}
