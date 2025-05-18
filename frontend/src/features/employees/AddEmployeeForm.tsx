import React, { useState, useEffect } from 'react';
import type { EmployeeCreate } from './employees.types';
import api from '../../api';

interface Position {
  id: number;
  name: string;
}

interface Props {
  onSuccess: () => void;
}

export default function AddEmployeeForm({ onSuccess }: Props) {
  const [positions, setPositions] = useState<Position[]>([]);
  const [form, setForm] = useState<EmployeeCreate>({
    name: '',
    surname: '',
    patronymic: '',
    sex: 'M',
    position_id: 0,
    salary: 0,
    birth_date: '',
    hired_at: '',
    child_number: 0,
  });

  useEffect(() => {
    api.get('/positions')
      .then(res => {
        const data = Array.isArray(res.data) ? res.data : res.data.items;
        setPositions(data.map((p: any) => ({ id: p.position_id, name: p.position_name })));
      });
  }, []);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm(f => ({
      ...f,
      [name]: name === 'salary' || name === 'child_number' || name === 'position_id'
        ? Number(value)
        : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (form.position_id === 0) {
      alert('Выберите должность');
      return;
    }
    try {
      await api.post('/employees', form);
      alert('Сотрудник добавлен');
      onSuccess();
      setForm({
        name: '',
        surname: '',
        patronymic: '',
        sex: 'M',
        position_id: 0,
        salary: 0,
        birth_date: '',
        hired_at: '',
        child_number: 0,
      });
    } catch (err) {
      alert('Ошибка при добавлении сотрудника');
      console.error(err);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input name="surname" value={form.surname} onChange={handleChange} placeholder="Фамилия" required />
      <input name="name" value={form.name} onChange={handleChange} placeholder="Имя" required />
      <input name="patronymic" value={form.patronymic} onChange={handleChange} placeholder="Отчество" />
      <select name="sex" value={form.sex} onChange={handleChange} required>
        <option value="M">М</option>
        <option value="F">Ж</option>
      </select>
      <select name="position_id" value={form.position_id} onChange={handleChange} required>
        <option value={0}>Выберите должность</option>
        {positions.map(p => (
          <option key={p.id} value={p.id}>{p.name}</option>
        ))}
      </select>
      <input name="salary" type="number" value={form.salary} onChange={handleChange} placeholder="Зарплата" required />
      <input name="birth_date" type="date" value={form.birth_date} onChange={handleChange} required />
      <input name="hired_at" type="date" value={form.hired_at} onChange={handleChange} required />
      <input name="child_number" type="number" value={form.child_number} onChange={handleChange} placeholder="Количество детей" required />
      <button type="submit">Добавить сотрудника</button>
    </form>
  );
}
