import './EmployeesTable.css';

import { useState, useEffect } from 'react';

const PassengerList = () => {
  const [passengers, setPassengers] = useState([]);
  const [filter, setFilter] = useState({
    routeID: '',
    travelDate: '',
    sex: '',
    minAge: '',
    maxAge: '',
    hasLuggage: '',
  });

  // Функция для обновления значений фильтров
  const handleFilterChange = (e) => {
    const { name, value } = e.target;
    setFilter((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  useEffect(() => {
    // Функция для получения списка пассажиров с фильтрами
    const fetchPassengers = async () => {
      try {
        const params = new URLSearchParams(filter);
        const response = await fetch(`http://localhost:8080/passengers/filter?${params.toString()}`);
        const data = await response.json();
        setPassengers(data);
      } catch (err) {
        console.error('Error fetching passengers:', err);
      }
    };

    fetchPassengers();
  }, [filter]);

  return (
    <div>
      <h2>Список Пассажиров</h2>

      {/* Фильтры */}
      <div>
        <label>Route ID:</label>
        <input
          type="text"
          name="routeID"
          value={filter.routeID}
          onChange={handleFilterChange}
        />
        <label>Дата поездки:</label>
        <input
          type="date"
          name="travelDate"
          value={filter.travelDate}
          onChange={handleFilterChange}
        />
        <label>Пол:</label>
        <select
          name="sex"
          value={filter.sex}
          onChange={handleFilterChange}
        >
          <option value="">Все</option>
          <option value="M">Мужской</option>
          <option value="F">Женский</option>
        </select>
        <label>Минимальный возраст:</label>
        <input
          type="number"
          name="minAge"
          value={filter.minAge}
          onChange={handleFilterChange}
        />
        <label>Максимальный возраст:</label>
        <input
          type="number"
          name="maxAge"
          value={filter.maxAge}
          onChange={handleFilterChange}
        />
        <label>Наличие багажа:</label>
        <select
          name="hasLuggage"
          value={filter.hasLuggage}
          onChange={handleFilterChange}
        >
          <option value="">Не важно</option>
          <option value="true">Есть</option>
          <option value="false">Нет</option>
        </select>
      </div>

      {/* Таблица пассажиров */}
      <table className="employees-table">
        <thead>
          <tr>
            <th>Имя</th>
            <th>Пол</th>
            <th>Возраст</th>
            <th>Номер маршрута</th>
            <th>Дата поездки</th>
            <th>Наличие багажа</th>
          </tr>
        </thead>
        <tbody>
          {passengers.map((passenger) => (
            <tr key={passenger.passenger_id}>
              <td>{passenger.name}</td>
              <td>{passenger.sex}</td>
              <td>{passenger.age}</td>
              <td>{passenger.route_id}</td>
              <td>{new Date(passenger.travel_date).toLocaleDateString()}</td>
              <td>{passenger.has_luggage ? 'Да' : 'Нет'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PassengerList;
