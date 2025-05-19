// src/pages/Home/Home.tsx
import { useState } from 'react';
import Tabs, { Tab } from '../../layouts/Tabs';
import Employees from '../../features/employees/EmployeesTab';
import Passengers      from '../../features/passengers/PassengersTab';
import Tickets         from '../../features/tickets/TicketsTab';
import Departments         from '../../features/departments/DepartmentsTab';

export default function Home() {
  const [activeTab, setActiveTab] = useState<Tab['id']>('employees');
  const tabs: Tab[] = [
    { id: 'employees',  label: 'Сотрудники' },
    { id: 'departments',  label: 'Департаменты' },
    { id: 'passengers', label: 'Пассажиры' },
    { id: 'tickets',    label: 'Билеты' },
  ];

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-4">Железнодорожная станция "Грасиона"</h1>
      <Tabs
        tabs={tabs}
        activeTab={activeTab}
        onChange={setActiveTab}
        children={{
          employees: <Employees />,
          departments: <Departments />,
          passengers: <Passengers />,
          tickets: <Tickets />,
        }}
      />
    </div>
  );
}
