package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
    var sum int32
    for _, num := range ar {
        sum += num
    }
    return sum
}

func main() {
    testCases := []struct {
        input    []int32
        expected int32
    }{
        {[]int32{1, 2, 3, 4, 10, 11}, 31},
        {[]int32{}, 0},                    
        {[]int32{-5, 0, 5}, 0},             
    }

    // Run and verify test cases
    for i, tc := range testCases {
        result := simpleArraySum(tc.input)
        status := "FAILED"
        if result == tc.expected {
            status = "PASSED"
        }
        fmt.Printf("Test case %d %s - Input: %v, Result: %d, Expected: %d\n",
            i+1,
            status,
            tc.input,
            result,
            tc.expected)
    }
}