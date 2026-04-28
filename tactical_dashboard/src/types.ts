export type SpectrumPoint = {
  frequency_mhz: number;
  power_db: number;
};

export type Burst = {
  start_index: number;
  end_index: number;
  peak_power_db: number;
};

export type Threat = {
  protocol: string;
  confidence: number;
  estimated_distance_m: number;
  rssi_dbm: number;
  last_seen: string;
  severity: string;
  reason: string;
  peaks: SpectrumPoint[];
};

export type Scan = {
  captured_at: string;
  center_frequency_mhz: number;
  bandwidth_mhz: number;
  noise_floor_db: number;
  rssi_dbm: number;
  spectrum: number[];
  peaks: SpectrumPoint[];
  bursts: Burst[];
  threats: Threat[];
};
