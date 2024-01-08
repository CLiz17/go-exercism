//Package Raindrops to convert a number into a string that contains raindrop sounds corresponding to certain potential factors.
package raindrops

import "strconv"

//Convert function is to convert a number into a string that contains raindrop sounds corresponding to factors 3, 5, 7
func Convert(number int) string {
	result := ""
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
