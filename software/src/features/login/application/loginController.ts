import React, { useState } from 'react';
import { useLoginHandler } from './loginHandler';
import type { LoginRequest } from '../entitys/loginEntity';
import type { ApiResult } from '../../../app/types';

const MAX_ATTEMPTS = 3;
const BLOCK_TIME = 30; // 30 seconds

export function useLoginController() {

    const { submit } = useLoginHandler();
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [attempts, setAttempts] = useState(0);
    const [blockedUntil, setBlockedUntil] = useState<number | null>(() => {
        const stored = localStorage.getItem('loginBlockedUntil');
        return stored ? Number(stored) : null;
    });
    const [blockSeconds, setBlockSeconds] = useState<number>(0);

    // Atualiza o contador de segundos do bloqueio
    React.useEffect(() => {
        if (blockedUntil && Date.now() < blockedUntil) {
            const updateSeconds = () => {
                const seconds = Math.max(0, Math.ceil((blockedUntil - Date.now()) / 1000));
                setBlockSeconds(seconds);
                if (seconds === 0) {
                    setBlockedUntil(null);
                    localStorage.removeItem('loginBlockedUntil');
                    setError(null);
                }
            };
            updateSeconds();
            const interval = setInterval(updateSeconds, 1000);
            return () => clearInterval(interval);
        } else {
            setBlockSeconds(0);
        }
    }, [blockedUntil]);

    async function onSubmit(input: LoginRequest): Promise<ApiResult<unknown>> {
        if (blockedUntil && Date.now() < blockedUntil) {
            setError(`Too many failed attempts. Please try again in ${blockSeconds} seconds.`);
            return { ok: false, error: 'Blocked due to too many attempts' };
        }
        setLoading(true);
        setError(null);
        const result = await submit(input);
        setLoading(false);
        if (!result.ok) {
            setAttempts((a) => a + 1);
            if (attempts + 1 >= MAX_ATTEMPTS) {
                const unblockTime = Date.now() + BLOCK_TIME * 1000;
                setBlockedUntil(unblockTime);
                localStorage.setItem('loginBlockedUntil', String(unblockTime));
                setAttempts(0);
                setError(`Too many failed attempts. Please try again in ${BLOCK_TIME} seconds.`);
                setTimeout(() => {
                    setBlockedUntil(null);
                    localStorage.removeItem('loginBlockedUntil');
                }, BLOCK_TIME * 1000);
                return result;
            }
            let msg = "An unexpected error occurred. Please try again.";
            if (result.error?.includes("connect") || result.error?.includes("Failed to fetch")) {
                msg = "Unable to connect to the server. Please check your connection or try again later.";
            } else if (result.error?.toLowerCase().includes("unauthorized") || result.error?.toLowerCase().includes("401")) {
                msg = "Invalid login or password.";
            } else if (result.error?.toLowerCase().includes("not found")) {
                msg = "User not found.";
            }
            setError(msg);
            return result;
        }
        setAttempts(0);
        return result;
    }

    return { onSubmit, loading, error, blockSeconds };
}