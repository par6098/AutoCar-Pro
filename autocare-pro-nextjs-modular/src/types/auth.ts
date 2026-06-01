export type UserRole = 'admin' | 'employee' | 'customer';

export type LoginRequest = {
  email: string;
  password: string;
  role: UserRole;
};

export type RegisterRequest = {
  name: string;
  email: string;
  phone: string;
  password: string;
  role: UserRole;
};