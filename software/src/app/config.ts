export const APP_CONFIG = {
    VITE_API_BASE_URL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000',
    TOKEN_STORAGE_KEY: 'auth_token',
    TOKEN_EXPIRES_AT_KEY: 'auth_token_expires_at'
}