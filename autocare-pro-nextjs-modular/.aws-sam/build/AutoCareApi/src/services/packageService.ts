import { USE_MOCKS } from '@/config/api';
import { httpGet } from '@/lib/httpClient';
import { ServicePackage } from '@/types/domain';
import { packages } from './mockData';
export const packageService = { list: () => USE_MOCKS ? Promise.resolve(packages) : httpGet<ServicePackage[]>('/packages') };
