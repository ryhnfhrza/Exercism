package booking

import "time"

import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05" 

    t, _ := time.Parse(layout,date) 

    return t
}

// => 2019-07-25 13:45:00 +0000 UTC

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05" 

    t, err := time.Parse(layout, date) 
    if err != nil{
        fmt.Println(err)
    }
    now := time.Now()

	hasPassedOrNot := t.Before(now) 
    
    return hasPassedOrNot
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05" 

    t, err := time.Parse(layout, date) 
    if err != nil{
        fmt.Println(err)
    }

    hour := t.Hour()

    if hour >= 12 && hour < 18{
        return true
    }else{
        return false
    }
    
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05" 

    t, err := time.Parse(layout, date) 
    if err != nil{
        fmt.Println(err)
    }

    day := t.Weekday()
    dayDate := t.Day()
    month := t.Month()
    year := t.Year()
    hour:= t.Hour()
    minute:= t.Minute()

    text := fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.", day,month,dayDate,year,hour,minute)

    return text
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	annivDate := time.Date(2026,9,15,0,0,0,0,time.UTC)

    return annivDate
}
