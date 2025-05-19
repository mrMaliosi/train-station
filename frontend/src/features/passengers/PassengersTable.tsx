import React, { ChangeEvent, useEffect, useState } from 'react';
import type { PassengerWithInfo } from './passengers.types';
import '../../styles/TableStyle.css';

interface Props {
  passengers: PassengerWithInfo[];
  onFilterChange: (f: any) => void;
  onDelete: (id: number) => void;
}

export default function PassengersTable({ passengers, onFilterChange, onDelete }: Props) {
  const [form, setForm] = useState({ routeID: '', sex: '', minAge: '', maxAge: '', hasLuggage: '', abroad: '', travelDate: '' });

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm(f => ({ ...f, [name]: value }));
  };

  useEffect(() => {
    const params: any = {};
    if (form.routeID)    params.routeID    = Number(form.routeID);
    if (form.sex)        params.sex        = form.sex;
    if (form.minAge)     params.minAge     = Number(form.minAge);
    if (form.maxAge)     params.maxAge     = Number(form.maxAge);
    if (form.hasLuggage) params.hasLuggage = form.hasLuggage === 'true';
    if (form.abroad)     params.abroad     = form.abroad === 'true';
    if (form.travelDate) params.travelDate = form.travelDate;
    onFilterChange(params);
  }, [form, onFilterChange]);

  return (
    <div className="table-container">
      <div className="filters">
        <input name="routeID" placeholder="Route ID" value={form.routeID} onChange={handleChange} />
        <select name="sex" value={form.sex} onChange={handleChange}>
          <option value="">Пол</option>
          <option value="M">M</option>
          <option value="F">F</option>
        </select>
        <input name="minAge" placeholder="Мин. возраст" value={form.minAge} onChange={handleChange} />
        <input name="maxAge" placeholder="Макс. возраст" value={form.maxAge} onChange={handleChange} />
        <select name="hasLuggage" value={form.hasLuggage} onChange={handleChange}>
          <option value="">Багаж?</option>
          <option value="true">Да</option>
          <option value="false">Нет</option>
        </select>
        <select name="abroad" value={form.abroad} onChange={handleChange}>
          <option value="">Заграница?</option>
          <option value="true">Да</option>
          <option value="false">Нет</option>
        </select>
        <input name="travelDate" type="date" value={form.travelDate} onChange={handleChange} />
      </div>
      <table className="employees-table">
        <thead>
          <tr>
            <th>#</th>
            <th>ФИО</th>
            <th>Пол</th>
            <th>Возраст</th>
            <th>Дата поездки</th>
            <th>Route ID</th>
            <th>Багаж</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {passengers.map((p, i) => (
            <tr key={p.passenger_id}>
              <td>{i + 1}</td>
              <td>{`${p.surname} ${p.name}${p.patronimic ? ' ' + p.patronimic : ''}`}</td>
              <td>{p.sex}</td>
              <td>{p.age}</td>
              <td>{p.travel_date}</td>
              <td>{p.route_id}</td>
              <td>{p.has_luggage ? 'Да' : 'Нет'}</td>
              <td>
                <button onClick={() => onDelete(p.passenger_id)}>Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}