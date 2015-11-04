package main
import "fmt"
func main(){

     var i int64 = 1
     var j int64 = 2
     var sum int64 = 0
     var k int64

     for k = 0; k < 4000000; k++ {
	 
     	 k = i + j
	 i = j
	 j = k

	 if i % 2 == 0 { sum += i }
	}

    fmt.Println(sum)
}