package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 1; (6*i-1)*(6*i-1) <= nb; i++ {
		if (6*i-1)*(6*i-1) < 0 {
			return false
		}
		if nb%(6*i-1) == 0 || nb%(6*i+1) == 0 {
			return false
		}
	}
	return true
}