package piscine

import "ft"

const N = 8

func solveNQueens(z []int, i int) {
	if i == N {
		for _, c := range z {
			ft.PrintRune(rune(c+1+'0'))
		}
		return
	}
	for j := 0; j < N; j++ {
		valid := true
		for k := 0; k < i; k++ {
			if z[k] == j || z[k]-j == i-k || z[k]-j == -(i-k) {
				valid = false
				break
			}
		}
		if valid {
			z[i] = j
			solveNQueens(z, i+1)
		}
	}
}

func EightQueens() {
	z := make([]int, N+1)
	z[N] = '\n'-1-'0'
	solveNQueens(z, 0)
}
