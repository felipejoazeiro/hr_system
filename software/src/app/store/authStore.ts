import {create} from 'zustand';
import { APP_CONFIG } from '../config'; 

type AuthState = {
    token: string | null;
    position: string | null;
    expiresAt: number | null;
    setSession: (token: string, position: string, expiresAt: number) => void;
    clear: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
    token: localStorage.getItem(APP_CONFIG.TOKEN_STORAGE_KEY),
    position: localStorage.getItem(APP_CONFIG.TOKEN_STORAGE_KEY) || null,
    expiresAt: Number(localStorage.getItem(APP_CONFIG.TOKEN_EXPIRES_AT_KEY)) || null,
    setSession: (token, position, expiresAt) => {
        localStorage.setItem(APP_CONFIG.TOKEN_STORAGE_KEY, token);
        localStorage.setItem(APP_CONFIG.TOKEN_EXPIRES_AT_KEY, expiresAt.toString());
        set({ token, position, expiresAt });
    },
    clear: () => {
        localStorage.removeItem(APP_CONFIG.TOKEN_STORAGE_KEY);
        localStorage.removeItem(APP_CONFIG.TOKEN_EXPIRES_AT_KEY);
        set({ token: null, position: null, expiresAt: null });
    },
}));