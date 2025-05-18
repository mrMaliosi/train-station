// src/api.ts
import axios from 'axios';

// В Vite все env-переменные берутся из import.meta.env, и их имена должны начинаться с VITE_
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:5173';

/** Основной экземпляр axios для всего приложения */
const api = axios.create({
  baseURL: API_URL,
  headers: { 'Content-Type': 'application/json' },
  timeout: 10_000,
});

export default api;
