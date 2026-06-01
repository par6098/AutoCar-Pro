import { USE_MOCKS } from '@/config/api';
import { httpGet } from '@/lib/httpClient';
import { Employee } from '@/types/domain';
import { employees } from './mockData';
export const employeeService = { list: () => USE_MOCKS ? Promise.resolve(employees) : httpGet<Employee[]>('/employees') };
