import axios from 'axios';

// Указываем базовый URL для нашего бэкенда
const API_URL = 'http://localhost:5000'; // Подставь правильный адрес

// Функции для работы с бригадами
export const getBrigadeEmployees = async () => {
  const response = await axios.get(`${API_URL}/brigades/employees`);
  return response.data;
};

export const getBrigadeEmployeesCount = async () => {
  const response = await axios.get(`${API_URL}/brigades/employees/count`);
  return response.data;
};

// Функции для работы с сотрудниками
export const getFilteredEmployees = async (filters) => {
  const response = await axios.get(`${API_URL}/employees`, { params: filters });
  return response.data;
};

export const getLocomotiveDrivers = async () => {
  const response = await axios.get(`${API_URL}/employees/locomotive-drivers`);
  return response.data;
};

// Функции для работы с локомотивами
export const getLocomotives = async () => {
  const response = await axios.get(`${API_URL}/locomotives`);
  return response.data;
};

export const getLocomotivesCount = async () => {
  const response = await axios.get(`${API_URL}/locomotives/count`);
  return response.data;
};

// Функции для работы с поездами
export const getTrains = async () => {
  const response = await axios.get(`${API_URL}/trains`);
  return response.data;
};

export const getTrainsCount = async () => {
  const response = await axios.get(`${API_URL}/trains/count`);
  return response.data;
};

// Функции для работы с маршрутами
export const getFilteredRoutes = async (filters) => {
  const response = await axios.get(`${API_URL}/routes/filter`, { params: filters });
  return response.data;
};

export const getReturnedTicketsDuringDelay = async () => {
  const response = await axios.get(`${API_URL}/routes/returned-tickets`);
  return response.data;
};

// Функции для работы с билетами
export const getSoldTickets = async () => {
  const response = await axios.get(`${API_URL}/tickets/sold`);
  return response.data;
};

export const getUnsoldTickets = async () => {
  const response = await axios.get(`${API_URL}/tickets/unsold`);
  return response.data;
};

export const getReturnedTicketsCount = async () => {
  const response = await axios.get(`${API_URL}/tickets/returned/count`);
  return response.data;
};

// Функции для работы с пассажирами
export const getFilteredPassengers = async (filters) => {
  const response = await axios.get(`${API_URL}/passengers/filter`, { params: filters });
  return response.data;
};
