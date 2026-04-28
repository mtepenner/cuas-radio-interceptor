import { Threat } from '../types';

type Props = {
  threats: Threat[];
};

export function ThreatAlerts({ threats }: Props) {
  return (
    <section className="panel stack-gap">
      <div className="section-heading">
        <span>Threat Alerts</span>
        <small>{threats.length} active signatures</small>
      </div>
      {threats.length === 0 ? (
        <div className="empty-state">No recognized drone control links in the current sweep.</div>
      ) : (
        threats.map((threat) => (
          <article key={`${threat.protocol}-${threat.last_seen}`} className={`threat-card severity-${threat.severity}`}>
            <header>
              <strong>{threat.protocol}</strong>
              <span>{Math.round(threat.confidence * 100)}% confidence</span>
            </header>
            <p>{threat.reason}</p>
            <div className="threat-meta">
              <span>{threat.rssi_dbm.toFixed(1)} dBm</span>
              <span>{Math.round(threat.estimated_distance_m)} m</span>
              <span>{new Date(threat.last_seen).toLocaleTimeString()}</span>
            </div>
          </article>
        ))
      )}
    </section>
  );
}
