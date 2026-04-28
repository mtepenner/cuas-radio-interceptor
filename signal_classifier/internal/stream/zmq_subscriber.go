package stream

import "time"

type SpectrumPoint struct {
	FrequencyMHz float64 `json:"frequency_mhz"`
	PowerDB      float64 `json:"power_db"`
}

type Frame struct {
	CapturedAt          time.Time       `json:"captured_at"`
	CenterFrequencyMHz  float64         `json:"center_frequency_mhz"`
	BandwidthMHz        float64         `json:"bandwidth_mhz"`
	RSSIDBm             float64         `json:"rssi_dbm"`
	NoiseFloorDB        float64         `json:"noise_floor_db"`
	HopIntervalsMs      []float64       `json:"hop_intervals_ms"`
	Peaks               []SpectrumPoint `json:"peaks"`
	ChannelActivityDB   []float64       `json:"channel_activity_db"`
	DominantTxPowerDBm  float64         `json:"dominant_tx_power_dbm"`
	DominantProtocolTag string          `json:"dominant_protocol_tag"`
}

type Subscriber struct {
	frames []Frame
	index  int
}

func NewSubscriber() *Subscriber {
	return &Subscriber{frames: seedFrames()}
}

func (s *Subscriber) Next() Frame {
	frame := s.frames[s.index%len(s.frames)]
	s.index++
	frame.CapturedAt = time.Now().UTC()
	return frame
}

func seedFrames() []Frame {
	return []Frame{
		{
			CenterFrequencyMHz: 2440,
			BandwidthMHz:       80,
			RSSIDBm:            -44.5,
			NoiseFloorDB:       -97.0,
			HopIntervalsMs:     []float64{8.1, 8.4, 7.8, 8.0, 8.2},
			Peaks: []SpectrumPoint{
				{FrequencyMHz: 2418.5, PowerDB: -47},
				{FrequencyMHz: 2431.2, PowerDB: -43},
				{FrequencyMHz: 2447.0, PowerDB: -45},
				{FrequencyMHz: 2462.7, PowerDB: -46},
			},
			ChannelActivityDB:   []float64{-95, -93, -89, -67, -58, -51, -49, -46, -45, -47, -52, -64, -81, -90, -94, -95},
			DominantTxPowerDBm:  29,
			DominantProtocolTag: "ocusync",
		},
		{
			CenterFrequencyMHz: 915,
			BandwidthMHz:       20,
			RSSIDBm:            -51.0,
			NoiseFloorDB:       -98.5,
			HopIntervalsMs:     []float64{4.0, 3.9, 4.2, 4.1, 4.0, 3.8},
			Peaks: []SpectrumPoint{
				{FrequencyMHz: 906.2, PowerDB: -55},
				{FrequencyMHz: 909.7, PowerDB: -50},
				{FrequencyMHz: 914.1, PowerDB: -49},
				{FrequencyMHz: 918.8, PowerDB: -53},
			},
			ChannelActivityDB:   []float64{-97, -95, -88, -71, -60, -56, -51, -49, -52, -55, -63, -78, -91, -95, -98, -99},
			DominantTxPowerDBm:  24,
			DominantProtocolTag: "expresslrs",
		},
		{
			CenterFrequencyMHz: 5785,
			BandwidthMHz:       40,
			RSSIDBm:            -63.0,
			NoiseFloorDB:       -101.0,
			HopIntervalsMs:     []float64{12.5, 13.0, 11.9, 12.2},
			Peaks: []SpectrumPoint{
				{FrequencyMHz: 5778.5, PowerDB: -68},
				{FrequencyMHz: 5784.0, PowerDB: -64},
				{FrequencyMHz: 5791.6, PowerDB: -66},
			},
			ChannelActivityDB:   []float64{-101, -100, -97, -92, -89, -83, -75, -71, -69, -73, -82, -90, -96, -100, -102, -103},
			DominantTxPowerDBm:  20,
			DominantProtocolTag: "unknown",
		},
	}
}
