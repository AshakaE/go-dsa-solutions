package samples

import (
	"fmt"
	"os"
	"github.com/AshakaE/go-dsa-solutions/algorithms"
)

func MiniMaxSumTest() {
    testCases := []struct {
        input       []int32
        expectedMin int64
        expectedMax int64
    }{
        {
            input:       []int32{1, 2, 3, 4, 5},
            expectedMin: 10,
            expectedMax: 14,
        },
        {
            input:       []int32{7, 3, 1, 5, 9},
            expectedMin: 16,
            expectedMax: 24,
        },
        {
            input:       []int32{5, 5, 5, 5, 5},
            expectedMin: 20,
            expectedMax: 20,
        },
        {
            input:       []int32{1000000000, 1500000000, 2000000000, 2000000000, 2147483647},
            expectedMin: 6500000000,
            expectedMax: 7647483647,
        },
    }

    for i, tc := range testCases {

        oldStdout := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        algorithms.MiniMaxSum(tc.input)

        w.Close()
        os.Stdout = oldStdout

        var min, max int64
        fmt.Fscanf(r, "%d %d", &min, &max)

        status := "FAILED"
        if min == tc.expectedMin && max == tc.expectedMax {
            status = "PASSED"
        }

        fmt.Printf("Test case %d %s - Input: %v\n", i+1, status, tc.input)
        fmt.Printf("  Got: %d (min), %d (max)\n", min, max)
        fmt.Printf("  Want: %d (min), %d (max)\n\n", tc.expectedMin, tc.expectedMax)
    }
}