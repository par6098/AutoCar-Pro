import { USE_MOCKS } from '@/config/api';
import { httpGet, httpPatch, httpPost } from '@/lib/httpClient';
import { Booking, CreateBookingRequest } from '@/types/domain';

const bookings: Booking[] = [];

export const bookingService = {
  //newbooking: (payload: Omit<Booking,'id'>) => USE_MOCKS ? Promise.resolve({ ...payload, id: `#${Date.now()}` }) : httpPost('/bookings', payload),
  customerList: (customerId: string) => USE_MOCKS ? Promise.resolve(bookings.filter(b => b.id === customerId)) : httpGet<Booking[]>(`/bookings/customer/${encodeURIComponent(customerId)}`),
  list: () => USE_MOCKS ? Promise.resolve(bookings) : httpGet<Booking[]>('/bookings'),
  create: (payload: Omit<CreateBookingRequest,'id'>) => httpPost('/bookings', payload),
  updateStatus: (id: string, status: Booking['status']) => USE_MOCKS ? Promise.resolve({ id, status }) : httpPatch(`/bookings/${encodeURIComponent(id)}`, { status })
};
