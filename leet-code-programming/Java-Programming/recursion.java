
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