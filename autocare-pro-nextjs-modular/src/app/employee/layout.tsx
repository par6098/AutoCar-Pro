'use client';

import AuthGuard from '@/components/auth/AuthGuard';

export default function EmployeeLayout({
  children,
}: {
  children: React.ReactNode;
}) {

  return (
    <AuthGuard
      allowedRoles={['EMPLOYEE']}
    >
      {children}
    </AuthGuard>
  );
}