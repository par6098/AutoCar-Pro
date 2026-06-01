import { Booking, DashboardSummary, Employee, MessageThread, ServicePackage, Transaction } from '@/types/domain';

export const bookings: Booking[] = [
  { id:'#1042', customerName:'Priya Sharma', phone:'9876543210', serviceName:'Premium Exterior Wash', vehicle:'Hyundai Creta', schedule:'Today 11am', assignedTo:'Ravi Kumar', amount:799, status:'Completed' },
  { id:'#1043', customerName:'Arjun Mehta', phone:'9812345670', serviceName:'Interior Deep Clean', vehicle:'Maruti Baleno', schedule:'Today 2pm', assignedTo:'Suresh Pal', amount:1199, status:'In Progress' },
  { id:'#1044', customerName:'Nisha Patel', phone:'9855671234', serviceName:'Pick & Drop + Full Wash', vehicle:'Toyota Innova', schedule:'Today 4pm', assignedTo:'Amit Singh', amount:1499, status:'Scheduled' },
  { id:'#1045', customerName:'Rohit Gupta', phone:'9890123456', serviceName:'Gold Detailing Package', vehicle:'MG XUV700', schedule:'Tomorrow 10am', amount:3499, status:'Pending' }
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
  todaysJobs: bookings
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
