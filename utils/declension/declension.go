package declension

func Declension(n int, one, few, many string) string {
	n = n % 100
	if n >= 11 && n <= 19 {
		return many
	}
	n = n % 10
	switch n {
	case 1:
		return one
	case 2, 3, 4:
		return few
	default:
		return many
	}
}
