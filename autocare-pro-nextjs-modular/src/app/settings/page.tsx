import AuthGuard from '@/components/auth/AuthGuard';
import { AppShell } from '@/components/layout/AppShell';
export default function Page(){return <AuthGuard allowedRoles={['ADMIN']}><AppShell initialPortal='admin' initialPage='settings' /></AuthGuard>;}
