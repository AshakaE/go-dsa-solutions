package algorithms

func DiagonalDifference(arr [][]int32) int32 {
	var forwardDiagonal int32
	var backwardDiagonal int32
	iter := 1
	for i := range arr {
		forwardDiagonal += arr[i][i]
		backwardDiagonal += arr[i][len(arr[i])-iter]
		iter += 1
	}
	absoluteDifference := forwardDiagonal - backwardDiagonal

	if absoluteDifference < 0 {
		absoluteDifference = -absoluteDifference
	}
	return absoluteDifference
}
