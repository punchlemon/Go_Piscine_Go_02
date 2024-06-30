package piscine

func Sqrt(nb int) int {
	var i int = 0
	for ;i*i < nb && i*i > -1; i++ {
	}
	if i*i == nb {
		return i
	}
	return 0
}