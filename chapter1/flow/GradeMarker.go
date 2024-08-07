package flow

func GetGrade(points float64) string {
	if points >= 80.00 {
		return "A"
	}
	if points >= 65 {
		return "B"
	}
	if points >= 30 && points <= 65 {
		return "C"
	}
	if points > 20 && points < 30 {
		return "D"
	}

	return "F"
}
