import { USE_MOCKS } from '@/config/api';
import { httpGet } from '@/lib/httpClient';
import { DashboardSummary } from '@/types/domain';
import { dashboardSummary } from './mockData';
export const dashboardService = { getSummary: () => USE_MOCKS ? Promise.resolve(dashboardSummary) : httpGet<DashboardSummary>('/dashboard/summary') };
