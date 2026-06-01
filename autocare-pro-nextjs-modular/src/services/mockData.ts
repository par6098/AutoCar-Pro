import { Booking,  DashboardSummary, Employee, MessageThread, ServicePackage, Transaction } from '@/types/domain';
import { BookingStatus } from '@/types/domain';

export const mockBookings: Booking[] = [
  {
    id: 'BK1001',
    customerName: 'Rahul Sharma',
    phone: '9876543210',
    serviceName: 'Premium Wash',
    vehicle: 'Hyundai Creta',
    bookingDate: '2026-05-21',
    bookingTime: '11:00 AM',
    assignedTo: 'Rohit',
    amount: 799,
    status: 'CONFIRMED' as BookingStatus,
  },
  {
    id: 'BK1002',
    customerName: 'Aman Verma',
    phone: '9811122233',
    serviceName: 'Gold Detailing',
    vehicle: 'Kia Seltos',
    bookingDate: '2026-05-21',
    bookingTime: '02:00 PM',
    assignedTo: 'Deepak',
    amount: 3499,
    status: 'IN_PROGRESS' as BookingStatus,
  },
  {
    id: 'BK1003',
    customerName: 'Neha Kapoor',
    phone: '9898989898',
    serviceName: 'Basic Wash',
    vehicle: 'Maruti Baleno',

    bookingDate: '2026-05-22',
    bookingTime: '11:00 AM',
    amount: 299,
    status: 'PENDING' as BookingStatus,
  },
  {
    id: 'BK1004',
    customerName: 'Sandeep Singh',
    phone: '9765432109',
    serviceName: 'Pick & Drop',
    vehicle: 'Toyota Fortuner',
    bookingDate: '2026-05-22',
    bookingTime: '04:00 PM',
    assignedTo: 'Vikas',
    amount: 299,
    status: 'CONFIRMED' as BookingStatus,
  },
  {
    id: 'BK1005',
    customerName: 'Priya Mehta',
    phone: '9900011122',
    serviceName: 'Premium Wash',
    vehicle: 'Honda City',
    bookingDate: '2026-05-23',
    bookingTime: '02:00 PM',
    assignedTo: 'Arjun',
    amount: 799,
    status: 'COMPLETED' as BookingStatus,
  },
  {
    id: 'BK1006',
    customerName: 'Karan Malhotra',
    phone: '9887766554',
    serviceName: 'Gold Detailing',
    vehicle: 'BMW X1',
    bookingDate: '2026-05-23',
    bookingTime: '11:00 AM',
    amount: 3499,
    status: 'PENDING' as BookingStatus,
  },
  {
    id: 'BK1007',
    customerName: 'Simran Kaur',
    phone: '9870011223',
    serviceName: 'Basic Wash',
    vehicle: 'Tata Nexon',
    bookingDate: '2026-05-24',
    bookingTime: '04:00 PM',
    assignedTo: 'Manoj',
    amount: 299,
    status: 'COMPLETED' as BookingStatus,
  },
  {
    id: 'BK1008',
    customerName: 'Aditya Jain',
    phone: '9818181818',
    serviceName: 'Premium Wash',
    vehicle: 'Mahindra XUV700',
    bookingDate: '2026-05-24',
    bookingTime: '02:00 PM',
    assignedTo: 'Ravi',
    amount: 799,
    status: 'CANCELLED' as BookingStatus,
  },
];

export const dashboardSummary: DashboardSummary = {
  bookingsThisMonth: 148,
  revenueThisMonth: 240000,
  activeEmployees: 24,
  averageRating: 4.8,
  weeklyRevenue: [
    { day:'Mon', amount:21000 }, { day:'Tue', amount:27000 }, { day:'Wed', amount:18000 },
    { day:'Thu', amount:33000 }, { day:'Fri', amount:38400 }, { day:'Sat', amount:29000 }, { day:'Sun', amount:23000 }
  ],
  todaysJobs: mockBookings.filter(b => b.bookingDate.startsWith('2026-05-21'))
};

export const packages: ServicePackage[] = [
  { id:'basic-wash', name:'Basic Wash', category:'Washing', price:299, duration:'30 min', vehicleType:'All', status:'Active', features:['Exterior rinse & dry','Wheel cleaning','Window wipe'] },
  { id:'premium-wash', name:'Premium Wash', category:'Washing', price:799, duration:'60 min', vehicleType:'All', status:'Active', featured:true, features:['Full exterior wash','Interior vacuum','Dashboard wipe','Tyre shine'] },
  { id:'gold-detailing', name:'Gold Detailing', category:'Detailing', price:3499, duration:'4 hrs', vehicleType:'Sedan/SUV', status:'Active', features:['Full exterior polish','Deep interior clean','Leather conditioning','Engine bay cleaning','Ceramic wax coat'] },
  { id:'ceramic-coating', name:'Ceramic Coating (5yr)', category:'Protection', price:8999, duration:'2 days', vehicleType:'Sedan/SUV', status:'Draft', features:['Paint correction','Ceramic coating','Warranty card'] }
];

export const employees: Employee[] = [
  { id:'emp-1', name:'Ravi Kumar', role:'Senior Detailer', shift:'8am - 5pm', checkIn:'7:52 AM', jobsToday:3, performance:'Excellent' },
  { id:'emp-2', name:'Suresh Pal', role:'Washer', shift:'10am - 7pm', checkIn:'9:58 AM', jobsToday:1, performance:'Good' },
  { id:'emp-3', name:'Amit Singh', role:'Driver', shift:'9am - 6pm', checkIn:'9:05 AM', jobsToday:2, performance:'Excellent' },
  { id:'emp-4', name:'Vikram Rao', role:'Driver', shift:'8am - 5pm', jobsToday:0, performance:'On Leave' }
];

export const messageThreads: MessageThread[] = [
  { id:'t1', customerName:'Priya Sharma', vehicle:'Honda City', lastMessage:'Thank you! Car looks amazing 🚗', unread:false, messages:[
    { direction:'in', text:'Hi, is my Creta ready?', time:'10:45 AM' },
    { direction:'out', text:'Yes Priya, your car is done! Please come anytime.', time:'10:47 AM' },
    { direction:'in', text:'Thank you! Car looks amazing 🚗', time:'10:52 AM' }
  ]},
  { id:'t2', customerName:'Arjun Mehta', vehicle:'Maruti Baleno', lastMessage:'Is my car ready yet?', unread:true, messages:[] }
];

export const transactions: Transaction[] = [
  { id:'txn-1', customerName:'Priya Sharma', amount:799, method:'UPI', status:'Completed' },
  { id:'txn-2', customerName:'Arjun Mehta', amount:1199, method:'Card', status:'Pending' },
  { id:'txn-3', customerName:'Sandeep Rana', amount:299, method:'Cash', status:'Completed' }
];
