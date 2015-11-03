package main
import "fmt"
func main(){
  
  var sum int = 0
  i := 0

  for i < 1000 {
	if i %3==0 || i % 5==0 { sum += i }
	i++
	}
	
	fmt.Println(sum)
}