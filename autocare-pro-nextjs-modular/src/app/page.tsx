'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

export default function Home() {

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

    if (role === 'ADMIN') {

      router.push('/dashboard');

    } else if (role === 'EMPLOYEE') {

      router.push('/employee/dashboard');

    } else {

      router.push('/dashboard');
    }

  }, [router]);

  return null;
}