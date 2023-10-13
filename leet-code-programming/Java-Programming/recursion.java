
// Online Java Compiler
// Use this editor to write, compile and run your Java code online


// Sum triangle from array
// https://www.geeksforgeeks.org/sum-triangle-from-array/

import java.util.*;
class HelloWorld {
    public static void main(String[] args) {
         int[] intArray = new int[]{ 1,2,3,4,5 };
         myMethod(intArray);
    }
    static void myMethod(int[] arr) {
     if(arr.length == 0){
         return;
      }
      int l = arr.length-1;
      int[] tempArr = new int[l];
      helper(tempArr, arr, 0);
      myMethod(tempArr);
      System.out.println(Arrays.toString(arr));
  }
    static int[] helper(int[] tempArr, int[] arr, int i ){
            if(i == tempArr.length){
                return tempArr;
            }
            tempArr[i] = arr[i]+arr[i+1];
            return helper(tempArr, arr, i+1);
    }
}

// https://www.geeksforgeeks.org/recursive-programs-to-find-minimum-and-maximum-elements-of-array/

// Input: arr = {1, 4, 3, -5, -4, 8, 6}; 
// Output: min = -5, max = 8 
// Input: arr = {1, 4, 45, 6, 10, -8}; 
// Output: min = -8, max = 45
import java.util.*;
class HelloWorld {
    public static void main(String[] args) {
        int[] arr = new int[]{1, 4, 45, 6, 10, -8}; 
        myMethod(arr, arr[0], arr[0], 0);
    }
    static void myMethod(int[] arr, int min, int max, int index) {
        if(index == arr.length){
            System.out.println("minimum "+ min+" maximun "+ max);
            return;
        }
        if(arr[index]<min){
            min = arr[index];
        }
        if(max<arr[index]){
            max = arr[index];
        }
        myMethod(arr, min, max, index+1);
    }
}