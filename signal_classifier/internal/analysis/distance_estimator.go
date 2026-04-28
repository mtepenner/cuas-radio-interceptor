package analysis

import "math"

func EstimateDistanceMeters(frequencyMHz, txPowerDBm, rxPowerDBm float64) float64 {
	if frequencyMHz <= 0 {
		return 0
	}

	pathLossDB := txPowerDBm - rxPowerDBm
	wavelengthMeters := 299792458.0 / (frequencyMHz * 1_000_000)
	return (wavelengthMeters / (4 * math.Pi)) * math.Pow(10, pathLossDB/20)
}
