'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

import { authService } from '@/services/authService';
import { UserRole } from '@/types/auth';

export default function LoginPage() {

  const router = useRouter();

  const [role, setRole] =
    useState<UserRole>('customer');

  const [email, setEmail] =
    useState('');

  const [password, setPassword] =
    useState('');

  const [loading, setLoading] =
    useState(false);

  const login = async () => {

  try {

    setLoading(true);

    const response =
      (await authService.login({
        email,
        password,
        role,
      })) as {
        access_token?: string;
        role?: UserRole;
      };

    console.log('LOGIN RESPONSE', response);

    /**
     * Expected backend response
     * {
     *   access_token: "...",
     *   role: "CUSTOMER"
     * }
     */

    const token =
      response?.access_token;

    if (!token) {
      throw new Error(
        'Access token missing'
      );
    }

    localStorage.setItem(
      'token',
      token
    );

    localStorage.setItem(
      'role',
      role
    );

    alert(`${role} login successful`);

    /**
     * Routing
     */

    if (role === 'admin') {

      router.push('/admin/dashboard');

    } else if (
      role === 'employee'
    ) {

      router.push(
        '/employee/dashboard'
      );

    } else {

      router.push('/dashboard');
    }

  } catch (error) {

    console.error(error);

    alert('Login failed');

  } finally {

    setLoading(false);
  }
};

  return (
    <div className="auth-page">

      <div className="card auth-card">

        <h2>Login</h2>

        <select
          className="form-select"
          value={role}
          onChange={(e) =>
            setRole(
              e.target.value as UserRole
            )
          }
        >
          <option value="ADMIN">
            Admin
          </option>

          <option value="EMPLOYEE">
            Employee
          </option>

          <option value="CUSTOMER">
            Customer
          </option>

        </select>

        <input
          className="form-input"
          placeholder="Email"
          value={email}
          onChange={(e) =>
            setEmail(e.target.value)
          }
        />

        <input
          className="form-input"
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) =>
            setPassword(e.target.value)
          }
        />

        <button
          className="btn btn-primary"
          onClick={login}
          disabled={loading}
        >
          {loading
            ? 'Logging in...'
            : 'Login'}
        </button>

        <button
          className="btn btn-ghost"
          onClick={() => router.push('/register')}
        >
          Register
        </button>

      </div>

    </div>
  );
}