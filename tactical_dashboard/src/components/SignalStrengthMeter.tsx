import { Threat } from '../types';

type Props = {
  primaryThreat: Threat | null;
  rssiDbm: number;
};

export function SignalStrengthMeter({ primaryThreat, rssiDbm }: Props) {
  const normalized = Math.max(0, Math.min(100, ((rssiDbm + 110) / 70) * 100));

  return (
    <section className="panel stack-gap">
      <div className="section-heading">
        <span>Signal Strength</span>
        <small>{primaryThreat ? primaryThreat.protocol : 'Unknown emitter'}</small>
      </div>
      <div className="meter-track">
        <div className="meter-fill" style={{ width: `${normalized}%` }} />
      </div>
      <div className="meter-meta">
        <span>{rssiDbm.toFixed(1)} dBm</span>
        <span>{primaryThreat ? `${Math.round(primaryThreat.estimated_distance_m)} m` : 'Range unavailable'}</span>
      </div>
    </section>
  );
}
