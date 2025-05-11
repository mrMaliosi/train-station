// App.jsx
import { useState } from 'react';
import './App.css';
import Tabs from './components/Tabs.jsx';
import BrigadeEmployees from './components/BrigadeEmployees';
import PassengerList from './components/PassengerList.jsx';

function App() {
  const [activeTab, setActiveTab] = useState('employees'); // Состояние для активной вкладки

  // Вкладки
  const tabs = [
    { id: 'employees', label: 'Сотрудники' },
    { id: 'other', label: 'Данные 2' }, // Дополнительная вкладка для примера
  ];

  // Функция для изменения активной вкладки
  const handleTabChange = (tabId) => {
    setActiveTab(tabId);
  };

  return (
    <>
      <div>
        <h2>Железнодорожная станция</h2>

        {/* Вкладки */}
        <div className="tabs">
          <button onClick={() => setActiveTab('employees')}>Сотрудники</button>
          <button onClick={() => setActiveTab('passengers')}>Пассажиры</button>
        </div>

        {/* Вывод активной вкладки */}
        {activeTab === 'employees' && <BrigadeEmployees />}
        {activeTab === 'passengers' && <PassengerList />}
      </div>
    </>
  );
}

export default App;
