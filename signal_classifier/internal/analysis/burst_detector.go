package analysis

type Burst struct {
	StartIndex  int     `json:"start_index"`
	EndIndex    int     `json:"end_index"`
	PeakPowerDB float64 `json:"peak_power_db"`
}

func DetectBursts(samples []float64, thresholdDB float64) []Burst {
	bursts := make([]Burst, 0)
	start := -1
	peak := thresholdDB

	flush := func(end int) {
		if start >= 0 {
			bursts = append(bursts, Burst{StartIndex: start, EndIndex: end, PeakPowerDB: peak})
			start = -1
			peak = thresholdDB
		}
	}

	for idx, power := range samples {
		if power >= thresholdDB {
			if start < 0 {
				start = idx
				peak = power
			}
			if power > peak {
				peak = power
			}
			continue
		}
		flush(idx - 1)
	}

	flush(len(samples) - 1)
	return bursts
}
