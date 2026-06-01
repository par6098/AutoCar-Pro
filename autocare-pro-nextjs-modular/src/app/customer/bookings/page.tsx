import AuthGuard from '@/components/auth/AuthGuard';
import { AppShell } from '@/components/layout/AppShell';
export default function Page(){return <AuthGuard allowedRoles={['CUSTOMER']}><AppShell initialPortal='customer' initialPage='customerBookings' /></AuthGuard>;}
