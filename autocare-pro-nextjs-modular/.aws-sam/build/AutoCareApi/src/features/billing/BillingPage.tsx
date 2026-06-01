'use client';
import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { Badge } from '@/components/ui/Badge';
import { billingService } from '@/services/billingService';
import { Transaction } from '@/types/domain';

export function BillingPage({ customerMode=false }: { customerMode?: boolean }) {
  const [rows, setRows] = useState<Transaction[]>([]);
  useEffect(() => { billingService.listTransactions().then(setRows); }, []);
  const total = rows.reduce((s, r) => s + r.amount, 0);
  return <><PageHeader title={customerMode ? 'Payments' : 'Billing & Payments'} subtitle="Invoices, payment gateway and transactions" actions={!customerMode && <button className="btn btn-accent">Create invoice</button>} />
  <div className="grid grid-2"><div className="card"><div className="card-title">Payment summary</div><div className="stat-val">₹{total}</div><div className="muted">Recent collected/pending value</div><button className="btn btn-primary" style={{marginTop:14}}>Collect payment</button></div><div className="card"><div className="card-title">Payment QR</div><div style={{display:'grid', placeItems:'center', height:140, border:'1px dashed var(--border)', borderRadius:12}}>UPI QR Placeholder</div></div></div>
  <div className="table-wrap"><table><thead><tr><th>ID</th><th>Customer</th><th>Amount</th><th>Method</th><th>Status</th></tr></thead><tbody>{rows.map(t => <tr key={t.id}><td className="mono">{t.id}</td><td>{t.customerName}</td><td className="mono">₹{t.amount}</td><td>{t.method}</td><td><Badge tone={t.status === 'Completed' ? 'green' : t.status === 'Pending' ? 'amber' : 'red'}>{t.status}</Badge></td></tr>)}</tbody></table></div></>;
}
