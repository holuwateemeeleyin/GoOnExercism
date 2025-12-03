package booking

import "time"

// Schedule returns a time.Time from a string containing a date (MM/DD/YYYY HH:MM:SS).
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed (format: "January 2, 2006 15:04:05").
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// (format: "Monday, January 2, 2006 15:04:05").
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment.
// Input: "7/25/2019 13:45:00"
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	t := Schedule(date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04") + "."
}

// AnniversaryDate returns this year's anniversary: September 15, current year at midnight.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
