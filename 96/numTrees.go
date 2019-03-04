package numTrees

func numTrees(n int) int {
	arr := make([]int, n+1)

	arr[0], arr[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i/2; j++ {
			arr[i] += 2 * (arr[j] * arr[i-1-j])
		}
		if i%2 == 1 {
			arr[i] += arr[i/2] * arr[i/2]
		}
	}

	return arr[n]
}
