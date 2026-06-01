'use client';
import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { Badge } from '@/components/ui/Badge';
import { employeeService } from '@/services/employeeService';
import { Employee } from '@/types/domain';

export function EmployeesPage() {
  const [rows, setRows] = useState<Employee[]>([]);
  useEffect(() => { employeeService.list().then(setRows); }, []);
  return <><PageHeader title="Team Management" subtitle="Employees, shifts, attendance and performance" actions={<button className="btn btn-accent">Add employee</button>} />
  <div className="grid grid-3" style={{marginBottom:14}}><div className="card"><div className="stat-val">{rows.length}</div><div className="muted">Total employees</div></div><div className="card"><div className="stat-val">{rows.filter(e=>e.performance !== 'On Leave').length}</div><div className="muted">On duty today</div></div><div className="card"><div className="stat-val">{rows.filter(e=>e.role.includes('Driver')).length}</div><div className="muted">Drivers</div></div></div>
  <div className="table-wrap"><table><thead><tr><th>Employee</th><th>Role</th><th>Shift</th><th>Check-in</th><th>Jobs</th><th>Performance</th></tr></thead><tbody>{rows.map(e => <tr key={e.id}><td><b>{e.name}</b></td><td>{e.role}</td><td>{e.shift}</td><td>{e.checkIn || '-'}</td><td>{e.jobsToday}</td><td><Badge tone={e.performance === 'Excellent' ? 'green' : e.performance === 'Good' ? 'blue' : e.performance === 'On Leave' ? 'gray' : 'amber'}>{e.performance}</Badge></td></tr>)}</tbody></table></div></>;
}
