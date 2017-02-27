#include <iostream>

int main(){
        
    long i, iterations = 0, max = 0, temp, targetNum;

    for(i = 2; i < 1000000; i++){
        
        temp = i;
        while(temp != 1){ //apply collatz sequence 
            if(temp % 2 == 0) //if even
                temp /= 2;
            else              //if odd
                temp = (temp * 3) + 1;

            iterations++;//track number of chains   
            }     

        if(max < iterations){ //find largest number of chains in seq
            max = iterations;
            targetNum = i; //store number with largest chains
        }
        iterations = 0;
    }

    std::cout << targetNum << std::endl;

    return 0;
}
