export type PageId = 'dashboard'|'bookings'|'packages'|'employees'|'messages'|'billing'|'settings'|'customerHome'|'customerBook'|'customerBookings'|'customerBilling'|'customerProfile';
export const adminNav = [
  { id:'dashboard', label:'Dashboard', section:'OVERVIEW', badge:undefined },
  { id:'bookings', label:'Bookings', section:'OVERVIEW', badge:12 },
  { id:'packages', label:'Services & Packages', section:'OPERATIONS' },
  { id:'employees', label:'Team', section:'OPERATIONS' },
  { id:'messages', label:'Messages', section:'BUSINESS', badge:5 },
  { id:'billing', label:'Billing', section:'BUSINESS' },
  { id:'settings', label:'Settings', section:'BUSINESS' }
] as const;
export const customerNav = [
  { id:'customerHome', label:'Home', section:'CUSTOMER' },
  { id:'customerBook', label:'Book a Service', section:'CUSTOMER' },
  { id:'customerBookings', label:'My Bookings', section:'CUSTOMER' },
  { id:'customerBilling', label:'Payments', section:'CUSTOMER' },
  { id:'customerProfile', label:'My Profile', section:'CUSTOMER' }
] as const;
