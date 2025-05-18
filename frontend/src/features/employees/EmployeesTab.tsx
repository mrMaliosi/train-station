// EmployeesTab.tsx
import React, { useEffect, useState } from 'react';
import type { Employee } from './employees.types';
import type { EmployeeFilters } from './employees.api';
import { getFilteredEmployees } from './employees.api';
import EmployeesTable from './EmployeesTable';
import AddEmployeeForm from './AddEmployeeForm';
import api from '../../api';

interface Department { id: number; name: string; }

export default function EmployeesTab() {
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [depts, setDepts] = useState<Department[]>([]);
  const [filters, setFilters] = useState<EmployeeFilters>({});

  useEffect(() => {
    api.get('/departments')
      .then(res => {
        const items = Array.isArray(res.data) ? res.data : res.data.items;
        setDepts(
          items.map((d: any) => ({ id: d.id, name: d.department_name }))
        );
      });
  }, []);

  useEffect(() => {
    getFilteredEmployees(filters)
      .then(setEmployees)
      .catch(console.error);
  }, [filters]);

  const refresh = () => getFilteredEmployees(filters).then(setEmployees);

  return (
    <>
      <AddEmployeeForm onSuccess={refresh} />
      <EmployeesTable
        employees={employees}
        departments={depts}
        onFilterChange={setFilters}
      />
    </>
  );
}