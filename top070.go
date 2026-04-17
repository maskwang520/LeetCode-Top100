package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	ppstep, pstep := 1, 1
	curSum := 0
	for i := 2; i <= n; i++ {
		curSum = ppstep + pstep
		ppstep = pstep
		pstep = curSum
	}
	return curSum
}
