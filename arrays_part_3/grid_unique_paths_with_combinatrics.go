package main

func uniquePaths1(m int, n int) int {
	N := m + n - 2 // (m-1+n-1) total no of steps to reach to bottom right
	r := m - 1     // to calculate Ncr
	res := 1
	//(N-r+1)(N-r+2)(N-r+3)...(N-r)
	//---------------------
	//1*2*3........r
	for i := 1; i <= r; i++ {
		res = res * (N - r + i) / i
	}
	return res
}
