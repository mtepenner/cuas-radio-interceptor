package protocol

import (
	"math"

	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/stream"
)

func MatchOcuSync(frame stream.Frame) (Match, bool) {
	if len(frame.HopIntervalsMs) < 4 || len(frame.Peaks) < 3 {
		return Match{}, false
	}

	mean := average(frame.HopIntervalsMs)
	if mean < 7.0 || mean > 9.5 {
		return Match{}, false
	}

	inBand := 0
	strongest := -200.0
	for _, peak := range frame.Peaks {
		if peak.FrequencyMHz >= 2400 && peak.FrequencyMHz <= 2485 {
			inBand++
		}
		strongest = math.Max(strongest, peak.PowerDB)
	}

	if inBand < 3 || strongest < -50 {
		return Match{}, false
	}

	confidence := clamp(0.62+0.08*float64(inBand)+0.01*(50+strongest), 0, 0.99)
	return Match{Protocol: "DJI OcuSync", Confidence: confidence, Reason: "8 ms cadence with dense 2.4 GHz multicarrier occupancy"}, true
}

func clamp(value, minValue, maxValue float64) float64 {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}
