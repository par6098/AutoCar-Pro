'use client';

import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { Badge } from '@/components/ui/Badge';
import { bookingService } from '@/services/bookingService';
import type { Booking } from '@/types/domain';

const tone = (status: Booking['status']) => {
  switch (status) {
    case 'COMPLETED':
      return 'green';
    case 'IN_PROGRESS':
      return 'blue';
    case 'CANCELLED':
      return 'red';
    case 'CONFIRMED':
    case 'SCHEDULED':
      return 'amber';
    case 'PENDING':
      return 'gray';
    default:
      return 'gray';
  }
};

export function BookingsPage({
  customerMode = true,
}: {
  customerMode?: boolean;
}) {
   const [rows, setRows] = useState<Booking[]>([]);

  useEffect(() => {
    alert('In bookingPage');
    bookingService.customerList('9749dbcf-bd88-4cab-89e1-b975214e9c1b').then(setRows);
  }, []);

  return (
    <>
      <PageHeader
        title={customerMode ? 'My Bookings' : 'All Bookings'}
        subtitle="Manage, assign and track service bookings"
        actions={
          <>
            <input className="form-input" placeholder="Search bookings..." />

            {!customerMode && (
              <button className="btn btn-accent" style={{ marginLeft: 8 }}>
                New booking
              </button>
            )}
          </>
        }
      />

      <div className="table-wrap">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Customer</th>
              <th>Service</th>
              <th>Vehicle</th>
              <th>Schedule</th>
              <th>Assigned</th>
              <th>Amount</th>
              <th>Status</th>
            </tr>
          </thead>

          <tbody>
            {rows.map((b) => (
              <tr key={b.id}>
                <td className="mono">{b.id}</td>

                <td>
                  <b>{b.customerName}</b>
                  <div className="muted">{b.phone}</div>
                </td>

                <td>{b.serviceName}</td>
                <td>{b.vehicle}</td>
                <td>{b.bookingTime  }</td>
                <td>{b.assignedTo || '-'}</td>
                <td className="mono">₹{b.amount}</td>

                <td>
                  <Badge tone={tone(b.status)}>{b.status}</Badge>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}