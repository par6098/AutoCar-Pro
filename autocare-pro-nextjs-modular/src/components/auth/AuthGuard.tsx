'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

type Props = {
  children: React.ReactNode;
  allowedRoles: string[];
};

export default function AuthGuard({
  children,
  allowedRoles,
}: Props) {

  const router = useRouter();

  useEffect(() => {

    const token =
      localStorage.getItem('token');

    const role =
      localStorage.getItem('role');

    if (!token) {

      router.push('/login');
      return;
    }

    if (!allowedRoles.includes(role || '')) {

      router.push('/login');
    }

  }, [router, allowedRoles]);

  return <>{children}</>;
}