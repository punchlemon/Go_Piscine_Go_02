package piscine

import "ft"

const N = 8

func Println(s string) {
	for _, c := range s {
		ft.PrintRune(c + 1)
	}
}

func printSolution(z []rune) {
	Println(string(z))
}

func solveNQueens(z []rune, x int, cnt *int) {
	if x == N {
		printSolution(z)
		*cnt++
		return
	}
	for i := 0; i < N; i++ {
		valid := true
		for j := 0; j < x; j++ {
			if z[j] == rune(i+'0') || z[j]-rune(i+'0') == rune(x-j) || z[j]-rune(i+'0') == rune(-(x-j)) {
				valid = false
				break
			}
		}
		if valid {
			z[x] = rune(i + '0')
			solveNQueens(z, x+1, cnt)
		}
	}
}

func EightQueens() {
	z := make([]rune, N+1)
	z[N] = '\n' - 1
	cnt := 0
	solveNQueens(z, 0, &cnt)
}