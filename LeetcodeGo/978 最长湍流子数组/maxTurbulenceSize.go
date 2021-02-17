package LeetCodeGo

func max(A int, B int) int {
	if A > B {
		return A
	}
	return B
}
func min(A int, B int) int {
	if A < B {
		return A
	}
	return B
}

// 动态规划方法
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	up := 1
	down := 1
	res := min(1, n)
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			up = down + 1
			down = 1
		} else if arr[i-1] > arr[i] {
			down = up + 1
			up = 1
		} else {
			down = 1
			up = 1
		}
		res = max(res, max(down, up))

	}
	return res

}
