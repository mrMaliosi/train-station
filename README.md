# train-station
Вариант: 4

## Языки
- [backend]: Go
- [frontend]: TypeScript

## Сборка и запуск
**1. Запустить бэкенд**
   ## При первом запуске:
   ```bash
   cd backend
   docker compose up --build
   ```

   *Если миграции не сработали при первом запуске, применить:
   ```bash
   ctrl+C
   docker compose up --build
   ```
   ## При повторных запусках:
   ```bash
   docker compose up
   ```

**2. Запустить фронтенд**
   ## При первом запуске:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

   ## При повторных запусках:
   ```bash
   npm run dev
   ```

## Формулировка задачи
### Цель 
Необходимо разработать клиент-серверное приложение железнодорожной станции.