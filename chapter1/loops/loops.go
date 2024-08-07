package loops

func Sum(a, b int) int {
	counter := a
	sum := 0
	for counter <= b {
		sum += counter
		counter++
	}
	return sum
}
