package analysis

import "testing"

func TestEstimateDistanceMeters(t *testing.T) {
	distance := EstimateDistanceMeters(2440, 28, -46)
	if distance < 45 || distance > 55 {
		t.Fatalf("expected realistic short-range estimate, got %.2f", distance)
	}
}
