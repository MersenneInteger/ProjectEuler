package main
import "fmt"
func main(){

     var i,j,sum, k int64 = 1, 2, 0, 0

     for k = 0; k < 4000000; k++ {
	 
     	 k = i + j
	 i = j
	 j = k

	 if i % 2 == 0 { sum += i }
	}

    fmt.Println(sum)
}