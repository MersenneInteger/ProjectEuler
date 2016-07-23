#include <stdio.h>
#include <stdbool.h>

bool isPalindrome(int array[], int size);
int getCount(int num);

int main(){
    
    int digit, count, product, copy, temp = 0, k = 0, max; 
    bool isPal = false;

    for(int i = 100; i < 1000; i++){
        for(int j = 999; j != i; j--){
           
            product = i * j;
            copy = product;
            count = getCount(copy);
            copy = product;
            int list[count];
            int counter = 0;    
            
            while(product != 0){

                digit = product % 10;
                list[counter] = digit;
                product /= 10;
                counter++;
            }

            isPal = isPalindrome(list, count);  
            if(isPal) {
                if(k == 0)
                    temp = copy;
                else
                    {
                        if(copy > temp) {
                                max = copy;
                                temp = copy;
                        }
                    }
                k++;
            }  
        }
    }//end of ounter for loop
    printf("%d ", max);

    return 0;        
}//end of main

bool isPalindrome(int array[], int size){
    bool flag = false;

    for(int i = 0; i < size; i++){
        if(array[i] != array[(size -i) - 1]){
            flag = false;
            break;
        }
        else {   
            flag = true;
            if(i == (size / 2)) break;
        }
    }//end of for loop

    if(flag) return true;
    else return false;

}//end of isPalindrome
int getCount(int num){
    int c = 0;

    while(num != 0){
        num /= 10;
        c++;
    }
    return c;
}//end of getCount
