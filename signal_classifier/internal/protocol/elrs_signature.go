package protocol

import (
	"math"

	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/analysis"
	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/stream"
)

type Match struct {
	Protocol   string  `json:"protocol"`
	Confidence float64 `json:"confidence"`
	Reason     string  `json:"reason"`
}

func MatchELRS(frame stream.Frame) (Match, bool) {
	if len(frame.HopIntervalsMs) < 4 || len(frame.Peaks) < 4 {
		return Match{}, false
	}

	mean := average(frame.HopIntervalsMs)
	if mean < 3.5 || mean > 4.5 {
		return Match{}, false
	}

	bursts := analysis.DetectBursts(frame.ChannelActivityDB, -62)
	if len(bursts) == 0 {
		return Match{}, false
	}

	spread := frame.Peaks[len(frame.Peaks)-1].FrequencyMHz - frame.Peaks[0].FrequencyMHz
	confidence := clamp(0.58+0.07*float64(len(bursts))+0.03*spread, 0, 0.98)
	return Match{Protocol: "ExpressLRS", Confidence: confidence, Reason: "4 ms hop cadence with broad sub-GHz burst spread"}, true
}

func average(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / math.Max(1, float64(len(values)))
}
