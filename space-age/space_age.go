package space

type Planet string
var secondsInOneEarthYear = 31557600.0
func Age(s float64, p Planet) float64 {
	m := map[Planet]float64{
		"Earth": 1.00,
		"Mercury": .2408467,
		"Venus": 0.61519726,
		"Mars":  1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}

	age := s / (m[p] * secondsInOneEarthYear)
	return age
}
