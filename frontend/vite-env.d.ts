interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  // добавьте сюда другие VITE_* переменные, если хотите
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
