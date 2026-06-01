'use client';
import { PageHeader } from '@/components/ui/PageHeader';
import { PageId } from '@/components/layout/navigation';

export function CustomerHomePage({ onNavigate }: { onNavigate: (page: PageId)=>void }) {
  return <><div className="card" style={{background:'var(--brand)', color:'white', padding:28}}><div style={{background:'var(--accent)', color:'var(--brand)', display:'inline-block', padding:'4px 10px', borderRadius:20, fontWeight:700}}>AUTOCARE PRO</div><h1 style={{marginTop:12}}>Premium car care at your doorstep</h1><p style={{opacity:.75, margin:'8px 0 18px'}}>Book washing, detailing, pick & drop and protection services.</p><button className="btn btn-accent" onClick={() => onNavigate('customerBook')}>Book service</button></div><PageHeader title="Popular services"/><div className="grid grid-4">{['Basic Wash','Premium Wash','Gold Detailing','Pick & Drop'].map((s,i)=><div className="card" key={s}><b>{s}</b><div className="muted">Starting ₹{[299,799,3499,299][i]}</div></div>)}</div></>;
}
export function CustomerBookPage() {
  return <><PageHeader title="Book a Service" subtitle="Select service, vehicle and preferred slot"/><div className="card"><div className="grid grid-2"><select className="form-select"><option>Premium Wash - ₹799</option><option>Gold Detailing - ₹3499</option></select><input className="form-input" placeholder="Vehicle e.g. Hyundai Creta"/><input className="form-input" type="date"/><select className="form-select"><option>11:00 AM</option><option>2:00 PM</option><option>4:00 PM</option></select></div><button className="btn btn-primary" style={{marginTop:14}}>Confirm booking</button></div></>;
}
