package main
import "fmt"
func main(){
     
	 var num, i, target int64 = 2520, 1, 0
     var isfound bool = false

     for !isfound {
     	for i = 1; i < 21; i++{ 
	        if num % i == 0 && num % 20 == 0 {
	            if i == 20 {        
                    isfound = true
	                target = num
                    break
                }
	        } else { 
                i = 1
		        num += 20	
            }
        }    
     }
     fmt.Print(target)
}
