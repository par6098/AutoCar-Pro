'use client';
import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { Badge } from '@/components/ui/Badge';
import { packageService } from '@/services/packageService';
import { ServicePackage } from '@/types/domain';

export function PackagesPage() {
  const [rows, setRows] = useState<ServicePackage[]>([]);
  useEffect(() => { packageService.list().then(setRows); }, []);
  return <><PageHeader title="Services & Packages" subtitle="Configure offerings, pricing and add-ons" actions={<button className="btn btn-accent">Add package</button>} />
  <div className="grid grid-3">{rows.map(p => <div key={p.id} className={`package ${p.featured ? 'featured' : ''}`}><div className="card-title">{p.name}<Badge tone={p.status === 'Active' ? 'green' : 'amber'}>{p.status}</Badge></div><div><span className="price">₹{p.price}</span><span className="muted"> / visit</span></div><div className="muted">{p.category} · {p.duration} · {p.vehicleType}</div><ul style={{marginTop:12, paddingLeft:18}}>{p.features.map(f => <li key={f}>{f}</li>)}</ul></div>)}</div></>;
}
