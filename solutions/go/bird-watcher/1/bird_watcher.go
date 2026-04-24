package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int
    
	for _,bird := range birdsPerDay{
        total = total + bird
    }

    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

    var start int
    var end int

	if week == 1 {
        start = 0
        end = 7
    }else if week == 2{
        start = 7
        end = 14
    }else if week == 3{
        start = 14
        end = 21
    }
    
	var total int
    
	for i := start; i < end; i++{
        total = total + birdsPerDay[i]
    }

    return total

    
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	logBirds := birdsPerDay[:]
    
	for i,_ := range logBirds {
        if i % 2 == 0 || i == 0{
            logBirds[i] = logBirds[i] +1
        }
    }

    return logBirds
}
