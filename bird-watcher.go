package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	l := len(birdsPerDay)
	for i := 0; i < l; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	for i := (week - 1) * 7; i < (week * 7); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

func FixBirdCountLog(birdsPerDay []int) []int {
	l := len(birdsPerDay)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
