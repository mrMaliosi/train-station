import React, { useEffect, useState } from 'react';
import type { DepartmentInfo } from './departments.types';
import { getDepartmentsInfo, deleteDepartment } from './departments.api';
import DepartmentsTable from './DepartmentsTable';

export default function DepartmentsTab() {
  const [departments, setDepartments] = useState<DepartmentInfo[]>([]);

  useEffect(() => {
    load();
  }, []);

  const load = () => {
    getDepartmentsInfo().then(setDepartments).catch(console.error);
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm('Удалить департамент?')) return;
    await deleteDepartment(id);
    load();
  };

  return (
    <>
      <DepartmentsTable
        departments={departments}
        onDelete={handleDelete}
      />
    </>
  );
}