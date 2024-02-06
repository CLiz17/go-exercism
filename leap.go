package leap

// IsLeapYear funtion is used to check if its a leap year or not
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if year%4 == 0 {
			return true
		} else {
			return false
		}
	}
}
