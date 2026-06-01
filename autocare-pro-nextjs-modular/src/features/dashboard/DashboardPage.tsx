'use client';
import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { Badge } from '@/components/ui/Badge';
import { dashboardService } from '@/services/dashboardService';
import { DashboardSummary } from '@/types/domain';
import { PageId } from '@/components/layout/navigation';

export function DashboardPage({ onNavigate }: { onNavigate: (page: PageId)=>void }) {
  const [data, setData] = useState<DashboardSummary>();
  useEffect(() => { dashboardService.getSummary().then(setData); }, []);
  if (!data) return <div>Loading dashboard...</div>;
  const max = Math.max(...data.weeklyRevenue.map(x => x.amount));
  return <>
    <PageHeader title="Dashboard" subtitle="Live operations overview" actions={<button className="btn btn-accent" onClick={() => onNavigate('bookings')}>New booking</button>} />
    <div className="grid grid-4">
      <Stat label="Bookings this month" value={data.bookingsThisMonth} />
      <Stat label="Revenue this month" value={`₹${Math.round(data.revenueThisMonth/100000)}.4L`} />
      <Stat label="Active employees" value={data.activeEmployees} />
      <Stat label="Average rating" value={data.averageRating} />
    </div>
    <div className="grid grid-2" style={{marginTop:14}}>
      <div className="card"><div className="card-title">Revenue - this week</div><div className="bars">{data.weeklyRevenue.map(x => <div key={x.day} style={{flex:1, textAlign:'center'}}><div className="bar" style={{height: `${Math.max(10, (x.amount/max)*100)}px`}}/><div className="muted">{x.day}</div></div>)}</div></div>
      <div className="card"><div className="card-title">Today's jobs</div>{data.todaysJobs.map(job => <div key={job.id} style={{display:'flex', justifyContent:'space-between', padding:'9px 0', borderBottom:'1px solid var(--border)'}}><div><b>{job.vehicle}</b><div className="muted">{job.serviceName} · {job.bookingTime}</div></div><Badge tone={job.status === 'Completed' ? 'green' : job.status === 'In Progress' ? 'blue' : 'amber'}>{job.status}</Badge></div>)}</div>
    </div>
  </>;
}
function Stat({label, value}:{label:string; value:string|number}) { return <div className="card"><div className="stat-val">{value}</div><div className="muted">{label}</div></div>; }
