/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_CALC_MIN?: string;
  readonly VITE_CALC_MAX?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
