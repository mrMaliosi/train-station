import './EmployeesTable.css';
import { useState, useEffect } from 'react';

const TicketInfo = () => {
  const [tickets, setTickets] = useState([]);
  const [statuses, setStatuses] = useState([]); // Статусы из ENUM
  const [soldCount, setSoldCount] = useState(0); // Количество проданных билетов
  const [returnedCount, setReturnedCount] = useState(0); // Количество возвращённых билетов
  const [filter, setFilter] = useState({
    routeID: '',
    fromDate: '',
    toDate: '',
    status: '',
  });

  const handleFilterChange = (e) => {
    const { name, value } = e.target;
    setFilter((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  // Получаем билеты
  useEffect(() => {
    const fetchTickets = async () => {
      try {
        const params = new URLSearchParams(filter);
        const response = await fetch(`http://localhost:8080/tickets?${params.toString()}`);
        const data = await response.json();
        setTickets(data);
      } catch (err) {
        console.error('Ошибка при получении билетов:', err);
      }
    };

    fetchTickets();
  }, [filter]);

  // Получаем статистику (количество проданных и возвращённых билетов)
  useEffect(() => {
  const fetchStats = async () => {
    try {
      const params = new URLSearchParams(filter);
      const response = await fetch(`http://localhost:8080/tickets/stats`);
      const data = await response.json();
      console.log('Received stats:', data); // Добавьте логирование для проверки
      setSoldCount(data.sold);
      setReturnedCount(data.returned);
    } catch (err) {
      console.error('Ошибка при получении статистики билетов:', err);
    }
  };

  fetchStats();
}, [filter]);


  // Получаем список возможных статусов
  useEffect(() => {
    const fetchStatuses = async () => {
      try {
        const res = await fetch(`http://localhost:8080/tickets/statuses`);
        const data = await res.json();
        setStatuses(data); // массив строк
      } catch (err) {
        console.error('Ошибка при получении статусов билетов:', err);
      }
    };

    fetchStatuses();
  }, []);

  // Функция для корректного отображения даты
  const formatDate = (date) => {
    if (date && date.Valid) {
      return new Date(date.Time).toLocaleDateString(); // Форматируем только если дата валидна
    }
    return '-'; // Если дата не валидна, показываем "-"
  };

  return (
    <div>
      <h2>Проданные билеты</h2>

      {/* Фильтры */}
      <div>
        <label>ID маршрута:</label>
        <input
          type="text"
          name="routeID"
          value={filter.routeID}
          onChange={handleFilterChange}
        />

        <label>С даты:</label>
        <input
          type="date"
          name="fromDate"
          value={filter.fromDate}
          onChange={handleFilterChange}
        />

        <label>По дату:</label>
        <input
          type="date"
          name="toDate"
          value={filter.toDate}
          onChange={handleFilterChange}
        />

        <label>Статус билета:</label>
        <select name="status" value={filter.status} onChange={handleFilterChange}>
          <option value="">Все</option>
          {statuses.map((status) => (
            <option key={status} value={status}>
              {status}
            </option>
          ))}
        </select>
      </div>

      {/* Статистика */}
      <div>
        <h3>Статистика</h3>
        <p>Продано билетов: {soldCount}</p>
        <p>Возвращено билетов: {returnedCount}</p>
      </div>

      {/* Таблица билетов */}
      <table className="employees-table">
        <thead>
          <tr>
            <th>ID билета</th>
            <th>ID маршрута</th>
            <th>Дата поездки</th>
            <th>Номер поезда</th>
            <th>Цена</th>
            <th>Статус</th>
          </tr>
        </thead>
        <tbody>
          {tickets.map((ticket) => (
            <tr key={ticket.TicketID}>
              <td>{ticket.TicketID}</td>
              <td>{ticket.RouteID}</td>
              <td>{formatDateTime(ticket.BoughtAt)}</td> {/* Используем formatDate для отображения даты */}
              <td>{ticket.TrainNumber}</td>
              <td>{ticket.Price}</td>
              <td>{ticket.TicketStatus}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TicketInfo;


// Функция для корректного отображения даты и времени
const formatDateTime = (date) => {
  if (date && date.Valid) {
    const dateTime = new Date(date.Time);
    const dateFormatted = dateTime.toLocaleDateString(); // Форматируем дату
    const timeFormatted = dateTime.toLocaleTimeString(); // Форматируем время
    return `${dateFormatted} ${timeFormatted}`; // Возвращаем строку с датой и временем
  }
  return '-'; // Если дата не валидна, показываем "-"
};
