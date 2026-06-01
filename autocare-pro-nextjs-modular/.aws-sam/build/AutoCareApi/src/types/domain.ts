export type BookingStatus = 'Pending' | 'Scheduled' | 'In Progress' | 'Completed' | 'Cancelled';

export interface DashboardSummary {
  bookingsThisMonth: number;
  revenueThisMonth: number;
  activeEmployees: number;
  averageRating: number;
  weeklyRevenue: { day: string; amount: number }[];
  todaysJobs: Booking[];
}

export interface Booking {
  id: string;
  customerName: string;
  phone: string;
  serviceName: string;
  vehicle: string;
  schedule: string;
  assignedTo?: string;
  amount: number;
  status: BookingStatus;
}

export interface ServicePackage {
  id: string;
  name: string;
  category: string;
  price: number;
  duration: string;
  vehicleType: string;
  status: 'Active' | 'Draft';
  features: string[];
  featured?: boolean;
}

export interface Employee {
  id: string;
  name: string;
  role: string;
  shift: string;
  checkIn?: string;
  jobsToday: number;
  performance: 'Excellent' | 'Good' | 'Average' | 'On Leave';
}

export interface MessageThread {
  id: string;
  customerName: string;
  vehicle: string;
  lastMessage: string;
  unread: boolean;
  messages: { direction: 'in' | 'out'; text: string; time: string }[];
}

export interface Transaction {
  id: string;
  customerName: string;
  amount: number;
  method: 'UPI' | 'Card' | 'Cash';
  status: 'Completed' | 'Pending' | 'Failed';
}
