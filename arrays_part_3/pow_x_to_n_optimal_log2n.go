package main

import "math"

/*
Algorithm
ans := 1
2^21
2*2^20 // if power is odd, take one no out to make power even, and multiply that no with ans variable
(2^2)^10 = 4^10, ans = 1*2
(4^2)^5 = 16^5, ans = 2
16^4, ans = 2*16
(16^2)^2, ans = 32
256^2, ans = 32
(256^2)^1, ans = 32
65536^1, ans = 32
65536^0, ans = 32 * 65536
*/
func myPow1(x float64, n int) float64 {
	nn := int(math.Abs(float64(n)))
	var xPowNn float64 = 1
	for nn > 0 {
		if nn%2 == 1 {
			xPowNn *= x
			nn -= 1
		} else {
			x *= x
			nn = nn / 2
		}
	}
	if n < 0 {
		return float64(1 / xPowNn)
	}
	return xPowNn
}
