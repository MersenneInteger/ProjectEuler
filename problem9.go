package main
import ("fmt" 
"math")
func main(){

    var squares [500]int
    for i:= 1; i < len(squares)+1; i++{
        squares[i-1] = i * i        
    }
    ans := findSum(squares)
    fmt.Print(ans)
}
func findSum(sqrList [500]int) int {
    
    for i:= 0; i < len(sqrList); i++{
        for j:= i+1; j < len(sqrList)-1; j++ {  
            temp := sqrList[i] + sqrList[j]
            for k:= j; k < len(sqrList); k++ {
                if temp == sqrList[k] {  
                    first := int(math.Sqrt(float64(sqrList[i])))
                    second := int(math.Sqrt(float64(sqrList[j])))
                    temp = int(math.Sqrt(float64(temp)))
                    sum := first + second + temp 
                    
                    if sum == 1000 {
                        return (first * second * temp)
                    }        
                }          
            }
        }        
    }
    return -1
}//end of findSum
