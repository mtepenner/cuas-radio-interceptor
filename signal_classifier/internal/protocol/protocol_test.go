package protocol

import (
	"testing"

	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/stream"
)

func TestMatchELRS(t *testing.T) {
	frame := stream.Frame{
		HopIntervalsMs: []float64{4.1, 4.0, 3.9, 4.2, 4.1},
		Peaks: []stream.SpectrumPoint{
			{FrequencyMHz: 905.5, PowerDB: -54},
			{FrequencyMHz: 910.1, PowerDB: -51},
			{FrequencyMHz: 914.7, PowerDB: -49},
			{FrequencyMHz: 918.9, PowerDB: -53},
		},
		ChannelActivityDB: []float64{-97, -81, -66, -55, -54, -73, -62, -56},
	}

	match, ok := MatchELRS(frame)
	if !ok || match.Protocol != "ExpressLRS" {
		t.Fatalf("expected ExpressLRS match, got %+v %v", match, ok)
	}
}

func TestMatchOcuSync(t *testing.T) {
	frame := stream.Frame{
		HopIntervalsMs: []float64{8.1, 7.9, 8.0, 8.3, 8.2},
		Peaks: []stream.SpectrumPoint{
			{FrequencyMHz: 2412.0, PowerDB: -48},
			{FrequencyMHz: 2437.0, PowerDB: -45},
			{FrequencyMHz: 2462.0, PowerDB: -46},
		},
	}

	match, ok := MatchOcuSync(frame)
	if !ok || match.Protocol != "DJI OcuSync" {
		t.Fatalf("expected OcuSync match, got %+v %v", match, ok)
	}
}
