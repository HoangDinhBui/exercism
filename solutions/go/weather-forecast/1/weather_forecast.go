// Package weather abc.
package weather

var (
    // CurrentCondition abc.
	CurrentCondition string
    // CurrentLocation abc.
	CurrentLocation  string
)

// Forecast abc.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
