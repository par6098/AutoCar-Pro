import AuthGuard from '@/components/auth/AuthGuard';
import { AppShell } from '@/components/layout/AppShell';
export default function Page(){return <AuthGuard allowedRoles={['ADMIN','CUSTOMER']}><AppShell initialPortal='admin' initialPage='dashboard' /></AuthGuard>;}
