// App.jsx
import { useState } from 'react';
import './App.css';
import Tabs from './components/Tabs.jsx';
import BrigadeEmployees from './components/BrigadeEmployees';
import PassengerList from './components/PassengerList.jsx';
import TicketInfo from './components/TicketInfo.jsx';

function App() {
  const [activeTab, setActiveTab] = useState('employees');

  // Добавляем все нужные вкладки
  const tabs = [
    { id: 'employees', label: 'Сотрудники' },
    { id: 'passengers', label: 'Пассажиры' },
    { id: 'tickets', label: 'Билеты' },
  ];

  return (
    <div>
      <h2>Железнодорожная станция</h2>

      {/* Отображение вкладок */}
      <div className="tabs">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id)}
            className={activeTab === tab.id ? 'active' : ''}
          >
            {tab.label}
          </button>
        ))}
      </div>

      {/* Отображение содержимого активной вкладки */}
      {activeTab === 'employees' && <BrigadeEmployees />}
      {activeTab === 'passengers' && <PassengerList />}
      {activeTab === 'tickets' && <TicketInfo />}
    </div>
  );
}

export default App;
