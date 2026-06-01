import { USE_MOCKS } from '@/config/api';
import { httpGet } from '@/lib/httpClient';
import { Transaction } from '@/types/domain';
import { transactions } from './mockData';
export const billingService = { listTransactions: () => USE_MOCKS ? Promise.resolve(transactions) : httpGet<Transaction[]>('/billing/transactions') };
