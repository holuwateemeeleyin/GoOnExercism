// Package weather provides functionality for describing and forecasting
// the current weather conditions in a given city.
package weather

var (
    // CurrentCondition stores the description of the weather condition.
	CurrentCondition string

    // CurrentLocation stores the name of the city.
	CurrentLocation  string
)

// Forecast returns a string describing the weather condition in the given city.
// It also updates the CurrentLocation and CurrentCondition variables.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
