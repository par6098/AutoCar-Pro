
'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService } from '@/services/authService';
import { UserRole } from '@/types/auth';

export default function RegisterPage() {
  const router = useRouter();
  const [role, setRole] = useState<UserRole>('customer');
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);

  const register = async () => {
    try {
      setLoading(true);
      await authService.register({ name, email, phone, password, role });
      alert(`${role} registered successfully`);
      router.push('/login');
    } catch (error) {
      console.error(error);
      alert('Registration failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-page">
      <div className="card auth-card">
        <h2>Register</h2>

        <select className="form-select" value={role} onChange={(e) => setRole(e.target.value as UserRole)}>
          <option value="ADMIN">Admin</option>
          <option value="EMPLOYEE">Employee</option>
          <option value="CUSTOMER">Customer</option>
        </select>

        <input className="form-input" placeholder="Full Name" value={name} onChange={(e) => setName(e.target.value)} />
        <input className="form-input" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
        <input className="form-input" placeholder="Phone" value={phone} onChange={(e) => setPhone(e.target.value)} />
        <input className="form-input" type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />

        <button className="btn btn-primary" onClick={register} disabled={loading}>
          {loading ? 'Registering...' : 'Register'}
        </button>

        <button className="btn btn-ghost" onClick={() => router.push('/login')}>
          Back to Login
        </button>
      </div>
    </div>
  );
}
