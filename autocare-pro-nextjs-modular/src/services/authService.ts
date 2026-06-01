import { httpPost } from '@/lib/httpClient';
import { LoginRequest, RegisterRequest } from '@/types/auth';

export const authService = {
  login(payload: LoginRequest) {
    return httpPost('/auth/login', payload);
  },

  register(payload: RegisterRequest) {
    return httpPost('/auth/register', payload);
  },
};