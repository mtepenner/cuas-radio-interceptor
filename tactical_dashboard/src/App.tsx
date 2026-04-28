import { useEffect, useMemo, useState } from 'react';

import { SignalStrengthMeter } from './components/SignalStrengthMeter';
import { ThreatAlerts } from './components/ThreatAlerts';
import { WaterfallDisplay } from './components/WaterfallDisplay';
import { Scan, Threat } from './types';

const apiBaseUrl = import.meta.env.VITE_CLASSIFIER_API_URL ?? 'http://127.0.0.1:8080';

export default function App() {
  const [scan, setScan] = useState<Scan | null>(null);
  const [threats, setThreats] = useState<Threat[]>([]);
  const [status, setStatus] = useState('Awaiting first sweep');

  useEffect(() => {
    let cancelled = false;

    const fetchData = async () => {
      try {
        const [scanResponse, threatsResponse] = await Promise.all([
          fetch(`${apiBaseUrl}/api/scan`),
          fetch(`${apiBaseUrl}/api/threats`),
        ]);

        if (!scanResponse.ok || !threatsResponse.ok) {
          throw new Error('Classifier API returned a non-200 response.');
        }

        const nextScan = (await scanResponse.json()) as Scan;
        const nextThreats = (await threatsResponse.json()) as Threat[];
        if (!cancelled) {
          setScan(nextScan);
          setThreats(nextThreats);
          setStatus(`Last sweep ${new Date(nextScan.captured_at).toLocaleTimeString()}`);
        }
      } catch (error) {
        if (!cancelled) {
          setStatus(error instanceof Error ? error.message : 'Unable to load classifier feed.');
        }
      }
    };

    fetchData();
    const intervalId = window.setInterval(fetchData, 1200);
    return () => {
      cancelled = true;
      window.clearInterval(intervalId);
    };
  }, []);

  const primaryThreat = useMemo(() => threats[0] ?? null, [threats]);

  return (
    <main className="shell">
      <section className="hero panel">
        <div>
          <p className="eyebrow">Counter-UAS RF Monitoring</p>
          <h1>Interception console for hopping control links and drone telemetry bursts.</h1>
        </div>
        <div className="hero-meta">
          <span>{scan ? `${scan.center_frequency_mhz.toFixed(0)} MHz` : 'Scanning bands'}</span>
          <span>{status}</span>
        </div>
      </section>

      <section className="dashboard-grid">
        <section className="panel stack-gap wide-panel">
          <div className="section-heading">
            <span>RF Waterfall</span>
            <small>{scan ? `${scan.bandwidth_mhz.toFixed(0)} MHz window` : 'Waiting for scan'}</small>
          </div>
          <WaterfallDisplay scan={scan} />
        </section>

        <ThreatAlerts threats={threats} />
        <SignalStrengthMeter primaryThreat={primaryThreat} rssiDbm={scan?.rssi_dbm ?? -110} />

        <section className="panel stack-gap">
          <div className="section-heading">
            <span>Spectrum Peaks</span>
            <small>{scan?.peaks.length ?? 0} dominant carriers</small>
          </div>
          <div className="peak-list">
            {(scan?.peaks ?? []).map((peak) => (
              <div key={`${peak.frequency_mhz}-${peak.power_db}`} className="peak-row">
                <span>{peak.frequency_mhz.toFixed(1)} MHz</span>
                <strong>{peak.power_db.toFixed(1)} dB</strong>
              </div>
            ))}
          </div>
        </section>
      </section>
    </main>
  );
}
