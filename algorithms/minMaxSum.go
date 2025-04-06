package algorithms

import (
	"fmt"
)

func MiniMaxSum(arr []int32) {

	// O(n^2) solution
	// var nums []int64  
    // ig := 0
    // for k := 0; k < len(arr); k++ {
    //     var num int64
    //     for i := range arr {
    //         if(i == ig){
    //             continue
    //         }
    //         num += int64(arr[i]) 
    //     }
    //     nums = append(nums, num)
    //      ig += 1
    // }
    // sort.Slice(nums, func(i, j int) bool {
	//     return nums[i] < nums[j] 
    // })
    // fmt.Println(nums[0], nums[len(nums) -1])

	// O(nlogn) solution
    // var min, max int64
    // slices.Sort(arr)
    // for i, value := range arr {
    //     if(i != 0) {
    //         max += int64(value)
    //     }
    //     if(i != len(arr) -1) {
    //         min += int64(value)
    //     }
    // }
    // fmt.Println(min, max)


	// O(n) solution
	var totalSum int64
    for _, num := range arr {
        totalSum += int64(num)
    }
    
    var minVal int64 = int64(arr[0])
    var maxVal int64 = int64(arr[0])
    
    for i := 1; i < len(arr); i++ {
        val := int64(arr[i])
        if val < minVal {
            minVal = val
        }
        if val > maxVal {
            maxVal = val
        }
    }
    fmt.Println(totalSum - maxVal, totalSum - minVal)
}