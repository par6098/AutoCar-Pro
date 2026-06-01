'use client';
import { useState } from 'react';
import { Sidebar } from './Sidebar';
import { PageId } from './navigation';
import { DashboardPage } from '@/features/dashboard/DashboardPage';
import { BookingsPage } from '@/features/bookings/BookingsPage';
import { PackagesPage } from '@/features/packages/PackagesPage';
import { EmployeesPage } from '@/features/employees/EmployeesPage';
import { MessagesPage } from '@/features/messages/MessagesPage';
import { BillingPage } from '@/features/billing/BillingPage';
import { CustomerHomePage, CustomerBookPage } from '@/features/customer/CustomerPages';

export function AppShell() {
  const [portal, setPortal] = useState<'admin'|'customer'>('admin');
  const [page, setPage] = useState<PageId>('dashboard');
  const switchPortal = (p: 'admin'|'customer') => { setPortal(p); setPage(p === 'admin' ? 'dashboard' : 'customerHome'); };
  const renderPage = () => ({
    dashboard:<DashboardPage onNavigate={setPage} />, bookings:<BookingsPage />, packages:<PackagesPage />, employees:<EmployeesPage />,
    messages:<MessagesPage />, billing:<BillingPage />, settings:<SettingsPage />,
    customerHome:<CustomerHomePage onNavigate={setPage} />, customerBook:<CustomerBookPage />, customerBookings:<BookingsPage customerMode />,
    customerBilling:<BillingPage customerMode />, customerProfile:<SettingsPage customerMode />
  }[page]);
  return <><header className="topbar"><div className="logo"><span className="logo-mark">A</span><span>AutoCare Pro</span></div><div className="portal-pill"><button className={portal === 'admin' ? 'active' : ''} onClick={() => switchPortal('admin')}>Admin</button><button className={portal === 'customer' ? 'active' : ''} onClick={() => switchPortal('customer')}>Customer</button></div></header><div className="shell"><Sidebar portal={portal} active={page} onNavigate={setPage}/><main className="main">{renderPage()}</main></div></>;
}
function SettingsPage({ customerMode=false }: { customerMode?: boolean }) { return <><div className="ph"><div><div className="ph-title">{customerMode?'My Profile':'Settings'}</div><div className="ph-sub">Business, payment, notification and access configuration</div></div></div><div className="card"><div className="card-title">{customerMode?'Profile information':'Business information'}</div><div className="grid grid-2"><input className="form-input" defaultValue={customerMode?'Priya Sharma':'AutoCare Pro Chandigarh'} /><input className="form-input" defaultValue={customerMode?'+91 98765 43210':'admin@autocarepro.com'} /></div><button className="btn btn-primary" style={{marginTop:14}}>Save changes</button></div></>; }
