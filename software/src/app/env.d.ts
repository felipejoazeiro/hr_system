interface ImportMetaEnv {
    readonly VITE_API_BASE_URL: string;
    readonly VITE_EXPIRES_AT_STORAGE_KEY: string;
    readonly VITE_POSITION_STORAGE_KEY: string;
    readonly VITE_TOKEN_STORAGE_KEY: string;
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}