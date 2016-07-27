package main
import ("fmt"
    "strconv")
func main(){

    var i, j, product, temp, max int64 = 0,0,0,0,0
    var isPal bool = false

    for i = 100; i < 1000; i++ {
        for j = 999; j != i; j-- {
                
                product = i * j
                isPal = isPalindrome(product)
                if isPal && product > temp {
                   max = product
                   temp = product
                }
        } 
    }
    fmt.Print(max) 

}//end of main
func isPalindrome(n int64) bool{
    
    str := strconv.FormatInt(int64(n), 10)
    size := len(str)
    for i := 0; i < size; i++ {
        
        if str[i] != str[(size -i) -1] {
            return false        
        }
        if i == (size/2) { break}
    }
    return true
}
