package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/analysis"
	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/protocol"
	"github.com/mtepenner/cuas-radio-interceptor/signal_classifier/internal/stream"
)

type threat struct {
	Protocol           string                 `json:"protocol"`
	Confidence         float64                `json:"confidence"`
	EstimatedDistanceM float64                `json:"estimated_distance_m"`
	RSSIDBm            float64                `json:"rssi_dbm"`
	LastSeen           time.Time              `json:"last_seen"`
	Severity           string                 `json:"severity"`
	Reason             string                 `json:"reason"`
	Peaks              []stream.SpectrumPoint `json:"peaks"`
}

type scan struct {
	CapturedAt         time.Time              `json:"captured_at"`
	CenterFrequencyMHz float64                `json:"center_frequency_mhz"`
	BandwidthMHz       float64                `json:"bandwidth_mhz"`
	NoiseFloorDB       float64                `json:"noise_floor_db"`
	RSSIDBm            float64                `json:"rssi_dbm"`
	Spectrum           []float64              `json:"spectrum"`
	Peaks              []stream.SpectrumPoint `json:"peaks"`
	Bursts             []analysis.Burst       `json:"bursts"`
	Threats            []threat               `json:"threats"`
}

type state struct {
	mu      sync.RWMutex
	latest  scan
	threats []threat
}

func main() {
	subscriber := stream.NewSubscriber()
	appState := &state{}

	go func() {
		ticker := time.NewTicker(900 * time.Millisecond)
		defer ticker.Stop()

		for {
			frame := subscriber.Next()
			bursts := analysis.DetectBursts(frame.ChannelActivityDB, -70)
			threats := classify(frame)

			appState.mu.Lock()
			appState.latest = scan{
				CapturedAt:         frame.CapturedAt,
				CenterFrequencyMHz: frame.CenterFrequencyMHz,
				BandwidthMHz:       frame.BandwidthMHz,
				NoiseFloorDB:       frame.NoiseFloorDB,
				RSSIDBm:            frame.RSSIDBm,
				Spectrum:           frame.ChannelActivityDB,
				Peaks:              frame.Peaks,
				Bursts:             bursts,
				Threats:            threats,
			}
			appState.threats = threats
			appState.mu.Unlock()

			<-ticker.C
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		respondJSON(writer, http.StatusOK, map[string]string{"service": "signal-classifier", "status": "ok"})
	})
	mux.HandleFunc("/api/scan", func(writer http.ResponseWriter, request *http.Request) {
		appState.mu.RLock()
		defer appState.mu.RUnlock()
		respondJSON(writer, http.StatusOK, appState.latest)
	})
	mux.HandleFunc("/api/threats", func(writer http.ResponseWriter, request *http.Request) {
		appState.mu.RLock()
		defer appState.mu.RUnlock()
		respondJSON(writer, http.StatusOK, appState.threats)
	})

	log.Println("signal classifier listening on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", cors(mux)); err != nil {
		log.Fatal(err)
	}
}

func classify(frame stream.Frame) []threat {
	matches := make([]threat, 0, 2)
	if match, ok := protocol.MatchOcuSync(frame); ok {
		matches = append(matches, buildThreat(frame, match))
	}
	if match, ok := protocol.MatchELRS(frame); ok {
		matches = append(matches, buildThreat(frame, match))
	}

	sort.Slice(matches, func(left, right int) bool {
		return matches[left].Confidence > matches[right].Confidence
	})
	return matches
}

func buildThreat(frame stream.Frame, match protocol.Match) threat {
	distance := analysis.EstimateDistanceMeters(frame.CenterFrequencyMHz, frame.DominantTxPowerDBm, frame.RSSIDBm)
	severity := "monitor"
	if match.Confidence >= 0.8 {
		severity = "critical"
	} else if match.Confidence >= 0.65 {
		severity = "high"
	}

	return threat{
		Protocol:           match.Protocol,
		Confidence:         match.Confidence,
		EstimatedDistanceM: distance,
		RSSIDBm:            frame.RSSIDBm,
		LastSeen:           frame.CapturedAt,
		Severity:           severity,
		Reason:             match.Reason,
		Peaks:              frame.Peaks,
	}
}

func respondJSON(writer http.ResponseWriter, status int, payload any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(payload); err != nil {
		log.Printf("encode error: %v", err)
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
