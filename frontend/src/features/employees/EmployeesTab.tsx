import React, { useEffect, useState } from 'react';
import type { Employee } from './employees.types';
import type { EmployeeFilters } from './employees.api';
import { getFilteredEmployees, deleteEmployee } from './employees.api';
import EmployeesTable from './EmployeesTable';
import AddEmployeeForm from './AddEmployeeForm';
import EditEmployeeModal from './EditEmployeeModal';
import api from '../../api';

interface Department { id: number; name: string; }

export default function EmployeesTab() {
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [depts, setDepts] = useState<Department[]>([]);
  const [filters, setFilters] = useState<EmployeeFilters>({});
  const [editEmp, setEditEmp] = useState<Employee | null>(null);
  const [isEditOpen, setIsEditOpen] = useState(false);

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
    refresh();
  }, [filters]);

  const refresh = () => getFilteredEmployees(filters).then(setEmployees).catch(console.error);

  const handleEdit = (emp: Employee) => {
    setEditEmp(emp);
    setIsEditOpen(true);
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Удалить сотрудника?')) {
      await deleteEmployee(id);
      refresh();
    }
  };

  const closeModal = () => setIsEditOpen(false);
  const onSave = () => refresh();

  return (
    <>
      <AddEmployeeForm onSuccess={refresh} />
      <EmployeesTable
        employees={employees}
        departments={depts}
        onFilterChange={setFilters}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />

      <EditEmployeeModal
        isOpen={isEditOpen}
        onClose={closeModal}
        employee={editEmp}
        onSave={onSave}
      />
    </>
  );
}