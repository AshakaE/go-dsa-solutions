package algorithms


func SimpleArraySum(ar []int32) int32 {
    var sum int32
    for _, num := range ar {
        sum += num
    }
    return sum
}
