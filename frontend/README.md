# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.



src/
│
├── api.ts                # HTTP-клиент, точки входа в бэкенд (RTK Query, axios и т.п.)
├── main.tsx              # точка монтирования React
├── index.css             # глобальные сбросы и утилиты Tailwind/Sass
│
├── assets/               # картинки, шрифты, статичные файлы
│
├── components/           # «размытые» общие компоненты, не привязанные к фиче
│   ├── Button/           # кнопка с вариантами (Primary, Secondary…)
│   │   ├── Button.tsx
│   │   └── Button.module.css
│   └── Modal/            # универсальный модальный диалог
│       └── Modal.tsx
│
├── layouts/              # «обёртки» и каркасы (Tabs, Header, Footer…)
│   └── Tabs.tsx
│
├── features/             # папки-фичи (каждая вкладка — своя фича)
│   ├── employees/
│   │   ├── EmployeesTab.tsx    # точка входа фичи (может быть index.tsx)
│   │   ├── EmployeesTable.tsx
│   │   ├── employees.api.ts     # вызовы к /employees
│   │   └── employees.types.ts   # TS-типы, константы
│   │
│   ├── passengers/
│   │   ├── PassengerList.tsx
│   │   ├── passengers.api.ts
│   │   └── passengers.types.ts
│   │
│   └── tickets/
│       ├── TicketInfo.tsx
│       ├── tickets.api.ts
│       └── tickets.types.ts
│
└── pages/                # «страницы» (у вас пока одна — Home)
    └── Home/
        ├── Home.tsx      # импортирует Tabs + подставляет контент
        └── Home.module.css
