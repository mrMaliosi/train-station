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
          value={filter.routeID}
          onChange={(e) => setFilter({ ...filter, routeID: e.target.value })}
        />
        <label>Дата поездки:</label>
        <input
          type="date"
          value={filter.travelDate}
          onChange={(e) => setFilter({ ...filter, travelDate: e.target.value })}
        />
        <label>Пол:</label>
        <select
          value={filter.sex}
          onChange={(e) => setFilter({ ...filter, sex: e.target.value })}
        >
          <option value="">Все</option>
          <option value="M">Мужской</option>
          <option value="F">Женский</option>
        </select>
        <label>Минимальный возраст:</label>
        <input
          type="number"
          value={filter.minAge}
          onChange={(e) => setFilter({ ...filter, minAge: e.target.value })}
        />
        <label>Максимальный возраст:</label>
        <input
          type="number"
          value={filter.maxAge}
          onChange={(e) => setFilter({ ...filter, maxAge: e.target.value })}
        />
        <label>Наличие багажа:</label>
        <select
          value={filter.hasLuggage}
          onChange={(e) => setFilter({ ...filter, hasLuggage: e.target.value })}
        >
          <option value="">Не важно</option>
          <option value="true">Есть</option>
          <option value="false">Нет</option>
        </select>
      </div>

      {/* Таблица пассажиров */}
      <table>
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
              <td>{passenger.travel_date}</td>
              <td>{passenger.has_luggage ? 'Да' : 'Нет'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PassengerList;
