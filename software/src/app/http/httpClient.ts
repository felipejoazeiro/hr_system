import { APP_CONFIG } from "../config";

const API_URL = APP_CONFIG.VITE_API_BASE_URL;

async function httpClient<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const config: RequestInit = {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...(options.headers || {})
        }
    };

    const response = await fetch(`${API_URL}${endpoint}`, config);

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    return response.json() as Promise<T>;
}

export async function get<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    return httpClient<T>(endpoint, { ...options, method: 'GET' });
}

export async function post<T>(endpoint: string, body: any, options: RequestInit = {}): Promise<T> {
    return httpClient<T>(endpoint, { ...options, method: 'POST', body: JSON.stringify(body) });
}

export async function put<T>(endpoint: string, body: any, options: RequestInit = {}): Promise<T> {
    return httpClient<T>(endpoint, { ...options, method: 'PUT', body: JSON.stringify(body) });
}

export async function patch<T>(endpoint: string, body: any, options: RequestInit = {}): Promise<T> {
    return httpClient<T>(endpoint, { ...options, method: 'PATCH', body: JSON.stringify(body) });
}