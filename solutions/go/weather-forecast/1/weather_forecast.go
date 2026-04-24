// Package weather ...
package weather

var (
    // CurrentCondition is variable type string .
	CurrentCondition string
    // CurrentLocation is variable type string.
	CurrentLocation  string
)

// Forecast ...
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
