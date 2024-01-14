package partyrobot

import (
	"fmt"
	"strconv"
)

func Welcome(name string) string {
	message := "Welcome to my party, " + name + "!"
	return message
}
func HappyBirthday(name string, age int) string {
	message := "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
	return message
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	formattedTable := fmt.Sprintf("%03d", table)
	formattedDistance := strconv.FormatFloat(distance, 'f', 1, 64)
	m1 := "Welcome to my party, " + name + "!\n"
	m2 := "You have been assigned to table " + formattedTable + ". Your table is " + direction + ", exactly " + formattedDistance + " meters from here.\n"
	m3 := "You will be sitting next to " + neighbor + "."

	return m1 + m2 + m3
}
