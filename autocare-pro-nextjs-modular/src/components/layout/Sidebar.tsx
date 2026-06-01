'use client';
import { PageId, adminNav, customerNav } from './navigation';
export function Sidebar({ portal, active, onNavigate }: { portal: 'admin'|'customer'; active: PageId; onNavigate: (id: PageId)=>void }) {
  const nav = portal === 'admin' ? adminNav : customerNav;
  const sections = [...new Set(nav.map(i => i.section))];
  return <aside className="sidebar">{sections.map(section => <div key={section}><div className="s-label">{section}</div>{nav.filter(i => i.section === section).map(item => <button key={item.id} className={`s-item ${active === item.id ? 'active' : ''}`} onClick={() => onNavigate(item.id as PageId)}><span>{item.label}</span>{'badge' in item && item.badge ? <span className="s-badge">{item.badge}</span> : null}</button>)}</div>)}<div className="card" style={{marginTop:20, boxShadow:'none'}}><b>AutoCare Pro</b><div className="muted">Chandigarh Branch</div><div className="muted">24 employees · Pro Plan</div></div></aside>;
}
