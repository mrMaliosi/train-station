// MainComponent.js
import React, { useState } from 'react';
import Tabs from './Tabs'; // Импортируем компонент вкладок
import EmployeesTable from './EmployeesTable'; // Ваш компонент с таблицей сотрудников

export default function DataTabs({ employees, otherData }) {
  const [activeTab, setActiveTab] = useState('employees'); // Хранение активной вкладки

  // Вкладки
  const tabs = [
    { id: 'employees', label: 'Сотрудники' },
    { id: 'other', label: 'Данные 2' },
  ];

  // Функция для смены активной вкладки
  const handleTabChange = (tabId) => {
    setActiveTab(tabId);
  };

  return (
    <div>
      {/* Вкладки */}
      <Tabs tabs={tabs} activeTab={activeTab} onTabChange={handleTabChange} />

      {/* Отображение контента в зависимости от выбранной вкладки */}
      {activeTab === 'employees' && <EmployeesTable employees={employees} />}
      {activeTab === 'other' && (
        <div>
          <h2>Данные 2</h2>
          <table className="other-data-table">
            <thead>
              <tr>
                <th>#</th>
                <th>Название</th>
                <th>Описание</th>
              </tr>
            </thead>
            <tbody>
              {otherData.map((item, idx) => (
                <tr key={idx}>
                  <td>{idx + 1}</td>
                  <td>{item.name}</td>
                  <td>{item.description}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
