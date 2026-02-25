package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (total int) {
	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekIndex := (week - 1) * 7
	birdsPerWeekDay := birdsPerDay[weekIndex:weekIndex+7]
	return TotalBirdCount(birdsPerWeekDay)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) (fixedBirdsPerDay []int) {
	fixedBirdsPerDay = birdsPerDay
	for i := 0; i < len(fixedBirdsPerDay); i += 2 {
		fixedBirdsPerDay[i]++
	}
	return
}
