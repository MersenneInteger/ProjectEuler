#include <stdio.h>
#include <stdbool.h>

bool isPalindrome(int array[], int size);
int getCount(int num);

int main(){
    
    int digit, count, product, copy, temp = 0, max; 
    bool isPal = false;

    for(int i = 100; i < 1000; i++){
        for(int j = 999; j != i; j--){
           
            copy = product = i * j; 
            count = getCount(copy);//get number of elements
            copy = product;
            int list[count], counter = 0;//create array to store digits    
            
            while(product != 0){ //breaks up digits

                digit = product % 10;
                list[counter] = digit;
                product /= 10;
                counter++;
            }

            isPal = isPalindrome(list, count); 
            if(isPal && copy > temp) {
                max = copy;
                temp = copy;   
            }
        }
    }//end of ounter for loop
    printf("%d ", max);

    return 0;        
}//end of main

bool isPalindrome(int array[], int size){
    
    bool flag = false;

    for(int i = 0; i < size; i++){//compare corresponding indexes 
        if(array[i] != array[(size -i) - 1])
            return false;
        else {   
            flag = true;
            if(i == (size / 2)) break;
        }
    }//end of for loop

     return true;

}//end of isPalindrome
int getCount(int num){
    
    int c = 0;

    while(num != 0){
        num /= 10;
        c++;
    }
    return c;
}//end of getCount
