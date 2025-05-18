// src/pages/Home/Home.tsx
import { useState } from 'react';
import Tabs, { Tab } from '../../layouts/Tabs';
import Employees from '../../features/employees/EmployeesTab';
//import PassengerList      from '../../features/passengers/PassengerList';
//import TicketInfo         from '../../features/tickets/TicketInfo';

export default function Home() {
  const [activeTab, setActiveTab] = useState<Tab['id']>('employees');
  const tabs: Tab[] = [
    { id: 'employees',  label: 'Сотрудники' },
    { id: 'passengers', label: 'Пассажиры' },
    { id: 'tickets',    label: 'Билеты' },
  ];

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-4">Железнодорожная станция</h1>
      <Tabs
        tabs={tabs}
        activeTab={activeTab}
        onChange={setActiveTab}
        children={{
          employees: <Employees />,
          passengers: <Employees />, //<PassengerList />,
          tickets: <Employees />, //<TicketInfo />,
        }}
      />
    </div>
  );
}
