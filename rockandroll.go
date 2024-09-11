package piscine

func RockAndRoll(n int) string {
	var m string
	if n >= 0 {
		if n%2 == 0 && n%3 != 0 {
			m = "rock\n"
		} else if n%3 == 0 && n%2 != 0 {
			m = "roll\n"
		} else if n%3 == 0 && n%2 == 0 {
			m = "rock and roll\n"
		} else {
			m = "error: non divisible\n"
		}
	} else if n < 0 {
		return "error: number is negative\n"
	}
	return m
}
