'use client';

import { useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { PageId } from '@/components/layout/navigation';
import { bookingService } from '@/services/bookingService';
import { Booking, CreateBookingRequest } from '@/types/domain';

export function CustomerHomePage({
  onNavigate,
}: {
  onNavigate: (page: PageId) => void;
}) {
  return (
    <>
      <div
        className="card"
        style={{
          background: 'var(--brand)',
          color: 'white',
          padding: 28,
        }}
      >
        <div
          style={{
            background: 'var(--accent)',
            color: 'var(--brand)',
            display: 'inline-block',
            padding: '4px 10px',
            borderRadius: 20,
            fontWeight: 700,
          }}
        >
          AUTOCARE PRO
        </div>

        <h1 style={{ marginTop: 12 }}>
          Premium car care at your doorstep
        </h1>

        <p style={{ opacity: 0.75, margin: '8px 0 18px' }}>
          Book washing, detailing, pick & drop and protection services.
        </p>

        <button
          className="btn btn-accent"
          onClick={() => onNavigate('customerBook')}
        >
          Book service
        </button>
      </div>

      <PageHeader title="Popular services" />

      <div className="grid grid-4">
        {['Basic Wash', 'Premium Wash', 'Gold Detailing', 'Pick & Drop'].map(
          (s, i) => (
            <div className="card" key={s}>
              <b>{s}</b>
              <div className="muted">
                Starting ₹{[299, 799, 3499, 299][i]}
              </div>
            </div>
          )
        )}
      </div>
    </>
  );
}

export function CustomerBookPage() {
  const [customerName, setCustomerName] = useState('');
  const [phone, setPhone] = useState('');
  const [serviceName, setServiceName] = useState('Premium Wash');
  const [vehicle, setVehicle] = useState('');
  const [bookingDate, setBookingDate] = useState('');
  const [bookingTime, setBookingTime] = useState('11:00 AM');
  const [loading, setLoading] = useState(false);

  const serviceAmounts: Record<string, number> = {
    'Basic Wash': 299,
    'Premium Wash': 799,
    'Gold Detailing': 3499,
    'Pick & Drop': 299,
  };

  const handleBooking = async () => {
    try {
      setLoading(true);

      const payload: CreateBookingRequest = {
        customer_id: '9749dbcf-bd88-4cab-89e1-b975214e9c1b', // In real app, this would come from auth context
        vehicle_id: '31e5f86e-70f7-4375-96ed-b76d1714f87e', // In real app, this would be selected from customer's vehicles
        service_package_id: '183a7543-e23c-4db4-aef0-deafd56fd9f0', // In real app, this would be selected based on serviceName
        booking_date: bookingDate,
        booking_slot: bookingTime,
        pickup_required: serviceName === 'Pick & Drop',
        pickup_address: serviceName === 'Pick & Drop' ? 'Customer provided address' : '',
      };

      await bookingService.create(payload);

      alert('Booking created successfully');

      // Reset form
      setCustomerName('');
      setPhone('');
      setVehicle('');
      setBookingDate('');
      setBookingTime('11:00 AM');
      setServiceName('Premium Wash');
    } catch (error) {
      console.error('Booking failed:', error);
      alert('Failed to create booking');
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <PageHeader
        title="Book a Service"
        subtitle="Select service, vehicle and preferred slot"
      />

      <div className="card">
        <div className="grid grid-2">
          <input
            className="form-input"
            placeholder="Customer Name"
            value={customerName}
            onChange={(e) => setCustomerName(e.target.value)}
          />

          <input
            className="form-input"
            placeholder="Phone Number"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
          />

          <select
            className="form-select"
            value={serviceName}
            onChange={(e) => setServiceName(e.target.value)}
          >
            <option value="Basic Wash">Basic Wash - ₹299</option>
            <option value="Premium Wash">Premium Wash - ₹799</option>
            <option value="Gold Detailing">Gold Detailing - ₹3499</option>
            <option value="Pick & Drop">Pick & Drop - ₹299</option>
          </select>

          <input
            className="form-input"
            placeholder="Vehicle e.g. Hyundai Creta"
            value={vehicle}
            onChange={(e) => setVehicle(e.target.value)}
          />

          <input
            className="form-input"
            type="date"
            value={bookingDate}
            onChange={(e) => setBookingDate(e.target.value)}
          />

          <select
            className="form-select"
            value={bookingTime}
            onChange={(e) => setBookingTime(e.target.value)}
          >
            <option>11:00 AM</option>
            <option>2:00 PM</option>
            <option>4:00 PM</option>
          </select>
        </div>

        <button
          className="btn btn-primary"
          style={{ marginTop: 14 }}
          onClick={handleBooking}
          disabled={loading}
        >
          {loading ? 'Booking...' : 'Confirm booking'}
        </button>
      </div>
    </>
  );
}