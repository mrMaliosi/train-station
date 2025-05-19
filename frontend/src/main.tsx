// src/main.tsx
import React from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';

// В новой структуре у нас одна «страница» — Home
import Home from './pages/Home/Home';

const container = document.getElementById('root');
if (!container) {
  throw new Error('Root container not found');
}

createRoot(container).render(
  <React.StrictMode>
    <Home />
  </React.StrictMode>
);
