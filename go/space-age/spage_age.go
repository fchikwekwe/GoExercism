package space

// Planet is a type that should only be a valid planet in our solar system as type string
type Planet string

// Age is a function to calculate how old someone is on different planets
func Age(seconds float64, planet Planet) float64 {
	var earthSeconds float64 = 31557600

	switch planet {
	case "Mercury":
		return seconds / (earthSeconds * 0.2408467)
	case "Venus":
		return seconds / (earthSeconds * 0.61519726)
	case "Mars":
		return seconds / (earthSeconds * 1.8808158)
	case "Jupiter":
		return seconds / (earthSeconds * 11.862615)
	case "Saturn":
		return seconds / (earthSeconds * 29.447498)
	case "Uranus":
		return seconds / (earthSeconds * 84.016846)
	case "Neptune":
		return seconds / (earthSeconds * 164.79132)
	}
	// If Earth is selected
	return seconds / (earthSeconds * 1)
}
