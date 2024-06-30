package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb < 2 {
		return 1
	} else {
		res := nb * RecursiveFactorial(nb-1)
		if res <= 0 {
			return 0
		}
		return res
	}
}