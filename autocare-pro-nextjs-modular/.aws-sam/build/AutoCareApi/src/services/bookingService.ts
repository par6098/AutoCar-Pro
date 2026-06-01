import { USE_MOCKS } from '@/config/api';
import { httpGet, httpPatch, httpPost } from '@/lib/httpClient';
import { Booking } from '@/types/domain';
import { bookings } from './mockData';
export const bookingService = {
  list: () => USE_MOCKS ? Promise.resolve(bookings) : httpGet<Booking[]>('/bookings'),
  create: (payload: Omit<Booking,'id'>) => USE_MOCKS ? Promise.resolve({ ...payload, id: `#${Date.now()}` }) : httpPost('/bookings', payload),
  updateStatus: (id: string, status: Booking['status']) => USE_MOCKS ? Promise.resolve({ id, status }) : httpPatch(`/bookings/${encodeURIComponent(id)}`, { status })
};
