import AuthGuard from '@/components/auth/AuthGuard';
import { AppShell } from '@/components/layout/AppShell';
export default function Page(){return <AuthGuard allowedRoles={['EMPLOYEE']}><AppShell initialPortal='admin' initialPage='dashboard' /></AuthGuard>;}
