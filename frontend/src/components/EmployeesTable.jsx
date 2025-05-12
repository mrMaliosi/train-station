import './EmployeesTable.css';

import React, { useState } from 'react';

export default function EmployeesTable({ employees }) {
  const [filteredEmployees, setFilteredEmployees] = useState(employees);
  const [filters, setFilters] = useState({
    position: '',
    sex: '',
    minSalary: '',
    maxSalary: '',
    teamID: '', // Новый фильтр для ID бригады
    minAge: '', // Новый фильтр для минимального возраста
    maxAge: '', // Новый фильтр для максимального возраста
  });

  // Функция для обновления значений фильтров
  const handleFilterChange = (e) => {
    const { name, value } = e.target;
    setFilters((prevFilters) => ({
      ...prevFilters,
      [name]: value,
    }));
  };

  // Функция для вычисления возраста сотрудника
  const calculateAge = (birthDate) => {
    const birth = new Date(birthDate);
    const now = new Date();
    let age = now.getFullYear() - birth.getFullYear();
    const m = now.getMonth() - birth.getMonth();
    if (m < 0 || (m === 0 && now.getDate() < birth.getDate())) {
      age--;
    }
    return age;
  };

  // Фильтрация данных
  const applyFilters = () => {
    const filtered = employees.filter((emp) => {
      const matchesPosition =
        filters.position === '' || emp.PositionName.includes(filters.position);
      const matchesSex = filters.sex === '' || emp.Sex === filters.sex;
      const matchesSalary =
        (filters.minSalary === '' || emp.Salary >= parseInt(filters.minSalary)) &&
        (filters.maxSalary === '' || emp.Salary <= parseInt(filters.maxSalary));
      const matchesTeamID =
        filters.teamID === '' || emp.TeamID.toString().includes(filters.teamID);
      const matchesAge =
        (filters.minAge === '' || calculateAge(emp.BirthDate.Time) >= parseInt(filters.minAge)) &&
        (filters.maxAge === '' || calculateAge(emp.BirthDate.Time) <= parseInt(filters.maxAge));

      return (
        matchesPosition &&
        matchesSex &&
        matchesSalary &&
        matchesTeamID &&
        matchesAge
      );
    });
    setFilteredEmployees(filtered);
  };

  // Вызываем фильтрацию при изменении фильтров
  React.useEffect(() => {
    applyFilters();
  }, [filters, employees]);

  return (
    <div className="table-container">
      <div className="filters">
        <input
          type="text"
          name="position"
          placeholder="Фильтр по должности"
          value={filters.position}
          onChange={handleFilterChange}
        />
        <select name="sex" value={filters.sex} onChange={handleFilterChange}>
          <option value="">Все</option>
          <option value="M">Мужской</option>
          <option value="F">Женский</option>
        </select>
        <input
          type="number"
          name="salary_from"
          placeholder="Мин. зарплата"
          value={filters.salary_from}
          onChange={handleFilterChange}
        />
        <input
          type="number"
          name="salary_to"
          placeholder="Макс. зарплата"
          value={filters.salary_to}
          onChange={handleFilterChange}
        />
        <input
          type="number"
          name="brigade_id"
          placeholder="Фильтр по ID бригады"
          value={filters.brigade_id}
          onChange={handleFilterChange}
        />
        <input
          type="number"
          name="age_from"
          placeholder="Мин. возраст"
          value={filters.age_from}
          onChange={handleFilterChange}
        />
        <input
          type="number"
          name="age_to"
          placeholder="Макс. возраст"
          value={filters.age_to}
          onChange={handleFilterChange}
        />
      </div>

      <table className="employees-table">
        <thead>
          <tr>
            <th>#</th>
            <th>ФИО</th>
            <th>Пол</th>
            <th>Должность</th>
            <th>Стаж</th>
            <th>Зарплата</th>
            <th>ID бригады</th> {/* Новый столбец */}
            <th>Возраст</th> {/* Новый столбец */}
          </tr>
        </thead>
        <tbody>
          {filteredEmployees.length === 0 ? (
            <tr>
              <td colSpan="8" className="no-data">
                Нет данных для отображения
              </td>
            </tr>
          ) : (
            filteredEmployees.map((emp, idx) => (
              <tr key={emp.ID}>
                <td>{idx + 1}</td>
                <td>{emp.Name} {emp.Surname} {emp.Patronymic}</td>
                <td>{emp.Sex === 'M' ? 'Мужской' : emp.Sex === 'F' ? 'Женский' : '-'}</td>
                <td>{emp.PositionName}</td>
                <td>{emp.Experience} лет</td>
                <td>{emp.Salary} ₽</td>
                <td>{emp.BrigadeID}</td> {/* Выводим ID бригады */}
                <td>{calculateAge(emp.BirthDate.Time)} лет</td> {/* Выводим возраст */}
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}

function getExperienceYears(hiredAt) {
  const hireDate = new Date(hiredAt.Time);
  const now = new Date();
  let years = now.getFullYear() - hireDate.getFullYear();
  const m = now.getMonth() - hireDate.getMonth();
  if (m < 0 || (m === 0 && now.getDate() < hireDate.getDate())) {
    years--;
  }
  return years;
}
