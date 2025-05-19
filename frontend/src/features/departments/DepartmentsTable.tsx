import React from 'react';
import type { DepartmentInfo } from './departments.types';
import '../../styles/TableStyle.css';

interface Props {
  departments: DepartmentInfo[];
  onDelete: (index: number) => void; // using index since no id returned
}

export default function DepartmentsTable({ departments, onDelete }: Props) {
  return (
    <div className="table-container">
      <table className="departments-table">
        <thead>
          <tr>
            <th>#</th>
            <th>Название</th>
            <th>Директор</th>
            <th>Дата рождения</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {departments.map((d, i) => (
            <tr key={i}>
              <td>{i + 1}</td>
              <td>{d.name}</td>
              <td>{
                d.directorName
                  ? `${d.directorSurname} ${d.directorName}${d.directorPatronymic ? ' ' + d.directorPatronymic : ''}`
                  : '-'
              }</td>
              <td>{d.birthDate || '-'}</td>
              <td>
                <button onClick={() => onDelete(i)}>Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}