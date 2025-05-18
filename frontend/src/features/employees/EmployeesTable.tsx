// EmployeesTable.tsx
import React, { useEffect, useState } from 'react';
import type { Employee, EmployeeFilterForm } from './employees.types';
import { calculateAge, getExperienceYears } from '../../utils/date';
import { getSexLabel } from '../../utils/data';
import '../../styles/TableStyle.css';

interface Props {
  employees: Employee[];
  departments: { id: number; name: string }[];
  onFilterChange: (apiFilters: any) => void;
}

export default function EmployeesTable({ employees, departments, onFilterChange }: Props) {
  const [form, setForm] = useState<EmployeeFilterForm>({
    departmentID: '', sex: '', ageFrom: '', ageTo: '',
    experienceFrom: '', experienceTo: '', childrenFrom: '', childrenTo: '',
    minSalary: '', maxSalary: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement|HTMLSelectElement>) => {
    const { name, value } = e.target as any;
    setForm(f => ({ ...f, [name]: value }));
  };
   const handleEdit = (id: number) => {
    console.log('Редактировать сотрудника с ID:', id);
    // TODO: Открыть форму редактирования
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm('Удалить этого сотрудника?')) return;
    try {
      //await api.delete(`/employees/${id}`);
      alert('Сотрудник удалён');
      onFilterChange({ ...form }); // повторный вызов фильтрации
    } catch (err) {
      alert('Ошибка при удалении');
      console.error(err);
    }
  };

  useEffect(() => {
    const api: any = {};
    if (form.departmentID)    api.departmentID    = +form.departmentID;
    if (form.sex)             api.sex             = form.sex;
    if (form.ageFrom)         api.ageFrom         = +form.ageFrom;
    if (form.ageTo)           api.ageTo           = +form.ageTo;
    if (form.experienceFrom)  api.experienceFrom  = +form.experienceFrom;
    if (form.experienceTo)    api.experienceTo    = +form.experienceTo;
    if (form.childrenFrom)    api.childrenFrom    = +form.childrenFrom;
    if (form.childrenTo)      api.childrenTo      = +form.childrenTo;
    if (form.minSalary)       api.minSalary       = +form.minSalary;
    if (form.maxSalary)       api.maxSalary       = +form.maxSalary;
    onFilterChange(api);
  }, [form]);

  return (
    <div className="table-container">
      <div className="filters">
        <select name="departmentID" value={form.departmentID} onChange={handleChange}>
          <option value="">Все отделы</option>
          {departments.map(d => <option key={d.id} value={d.id}>{d.name}</option>)}
        </select>
        <select name="sex" value={form.sex} onChange={handleChange}>
          <option value="">Пол (все)</option>
          <option value="M">М</option>
          <option value="F">Ж</option>
        </select>
        <input name="ageFrom" placeholder="Мин. возраст" value={form.ageFrom} onChange={handleChange} />
        <input name="ageTo"   placeholder="Макс. возраст" value={form.ageTo}   onChange={handleChange} />
        <input name="experienceFrom" placeholder="Мин. стаж" value={form.experienceFrom} onChange={handleChange} />
        <input name="experienceTo"   placeholder="Макс. стаж" value={form.experienceTo}   onChange={handleChange} />
        <input name="childrenFrom" placeholder="Мин. детей" value={form.childrenFrom} onChange={handleChange} />
        <input name="childrenTo"   placeholder="Макс. детей" value={form.childrenTo}   onChange={handleChange} />
        <input name="minSalary" placeholder="Мин. зарплата" value={form.minSalary} onChange={handleChange} />
        <input name="maxSalary" placeholder="Макс. зарплата" value={form.maxSalary} onChange={handleChange} />
      </div>

            <table className="employees-table">
        <thead>
          <tr>
            <th>#</th>
            <th>ФИО</th>
            <th>Пол</th>
            <th>Отдел</th>
            <th>Должность</th>
            <th>Стаж</th>
            <th>Зарплата</th>
            <th>Возраст</th>
            <th>Дети</th>
            <th>Действия</th> {/* Новый столбец */}
          </tr>
        </thead>
        <tbody>
          {employees.map((emp, i) => (
            <tr key={emp.ID}>
              <td>{i + 1}</td>
              <td>{`${emp.Surname} ${emp.Name}${emp.Patronymic ? ' ' + emp.Patronymic : ''}`}</td>
              <td>{getSexLabel(emp.Sex)}</td>
              <td>{emp.DepartmentName}</td>
              <td>{emp.PositionName}</td>
              <td>{emp.Experience ?? getExperienceYears(emp.HiredAt?.Time ?? emp.BirthDate.Time)}</td>
              <td>{emp.Salary} ₽</td>
              <td>{calculateAge(emp.BirthDate.Time)}</td>
              <td>{emp.ChildNumber ?? 0}</td>
              <td>
                <button onClick={() => handleEdit(emp.ID)}>Изменить</button>
                <button onClick={() => handleDelete(emp.ID)}>Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

    </div>
  );
}