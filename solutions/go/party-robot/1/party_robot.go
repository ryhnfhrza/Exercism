package partyrobot

import "strconv"
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name +"!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	
    
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	tableNumber := fmt.Sprintf("%03d",table)
    distanceNumber := fmt.Sprintf("%.1f", distance)

    // text := "Welcome to my party, " + name + "!\nYou have been assigned to table " + tableNumber + ". Your table is " + direction + ", exactly " + distanceNumber + " meters from here.\nYou will be sitting next to " + neighbor + "."

	text := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.",name,tableNumber,direction,distanceNumber,neighbor)
    
    return text;
}
