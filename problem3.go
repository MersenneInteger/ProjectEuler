/*The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main
import "fmt"
func main(){

var num, i, j, k, n int64 = 600851475143, 0, 0, 0, 0
var max[5] int64

for j = 2; j<= num/3; j++ {
    i = 2;
    for ; i <= j-1; i++ {
        if j%i == 0 {break}
    }
    if i == j && i!=2 {
       	   if num % j == 0 {
	      max[k] = j
	      k++
	      if k > 3 {break}
   
      }
    }
  }

for a:=0; a < len(max); a++ {
    if max[a] > n {
       n = max[a]
       }
     }
     fmt.Println(n)
}

