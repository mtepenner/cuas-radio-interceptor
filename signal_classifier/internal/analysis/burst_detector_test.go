package analysis

import "testing"

func TestDetectBursts(t *testing.T) {
	samples := []float64{-95, -82, -70, -68, -91, -94, -72, -71, -97}
	bursts := DetectBursts(samples, -75)

	if len(bursts) != 2 {
		t.Fatalf("expected 2 bursts, got %d", len(bursts))
	}

	if bursts[0].StartIndex != 2 || bursts[0].EndIndex != 3 {
		t.Fatalf("unexpected first burst bounds: %+v", bursts[0])
	}

	if bursts[1].PeakPowerDB != -71 {
		t.Fatalf("unexpected second burst peak: %+v", bursts[1])
	}
}
