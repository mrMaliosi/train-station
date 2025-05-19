import React, { ChangeEvent, useEffect, useState } from 'react';
import type { Ticket } from './tickets.types';
import '../../styles/TableStyle.css';

interface Props {
  tickets: Ticket[];
  onFilterChange: (f: any) => void;
  onDelete: (id: number) => void;
}

export default function TicketsTable({ tickets, onFilterChange, onDelete }: Props) {
  const [form, setForm] = useState({ routeID: '', fromDate: '', toDate: '', status: '' });

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm(f => ({ ...f, [name]: value }));
  };

  useEffect(() => {
    const params: any = {};
    if (form.routeID)  params.routeID  = Number(form.routeID);
    if (form.fromDate) params.fromDate = form.fromDate;
    if (form.toDate)   params.toDate   = form.toDate;
    if (form.status)   params.status   = form.status;
    onFilterChange(params);
  }, [form, onFilterChange]);

  return (
    <div className="table-container">
      <div className="filters">
        <input name="routeID" placeholder="Route ID" value={form.routeID} onChange={handleChange} />
        <input name="fromDate" type="date" placeholder="С" value={form.fromDate} onChange={handleChange} />
        <input name="toDate" type="date" placeholder="По" value={form.toDate} onChange={handleChange} />
        <input name="status" placeholder="Статус" value={form.status} onChange={handleChange} />
      </div>
      <table className="employees-table">
        <thead>
          <tr>
            <th>#</th>
            <th>Ticket Number</th>
            <th>Route ID</th>
            <th>Passenger ID</th>
            <th>Status</th>
            <th>Bought At</th>
            <th>Price</th>
            <th>Train #</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {tickets.map((t, i) => (
            <tr key={t.ticketId}>
              <td>{i + 1}</td>
              <td>{t.ticketId}</td>
              <td>{t.routeId}</td>
              <td>{t.passengerId}</td>
              <td>{t.status}</td>
              <td>{t.boughtAt ?? '-'}</td>
              <td>{t.price}</td>
              <td>{t.trainNumber}</td>
              <td>
                <button onClick={() => onDelete(t.ticketId)}>Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}