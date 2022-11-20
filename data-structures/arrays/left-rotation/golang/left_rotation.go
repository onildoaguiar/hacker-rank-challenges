package main

import (
	"fmt"
)

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func rotateLeft(d int32, arr []int32) []int32 {
    var count = d
    var rotatedArr = arr
    var arrLength = len(rotatedArr)
    
    for ok := true; ok; ok = ( count > 0) { 
        var tempArr = make([]int32, arrLength)
        var first int32
        
        for i := 0; i < arrLength; i++ {
            if i == 0 {
                first = rotatedArr[i]
            } else {
                tempArr[i-1] = rotatedArr[i]
            }
        }
        
        tempArr[arrLength-1] = first
        rotatedArr = tempArr
    
        count -= 1;
    }  
    
    return rotatedArr
}

func main() {
    intputArr := []int32 {1, 2, 3, 4, 5}
    var inputD int32 = 54 

    outputArr := rotateLeft(inputD, intputArr)
    fmt.Printf("outputArr: %v, %T\n", outputArr, outputArr)
}