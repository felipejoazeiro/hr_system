import { useState } from 'react';
import { useLoginHandler } from './loginHandler';
import type { LoginRequest } from '../entitys/loginEntity';
import type { ApiResult } from '../../../app/types';

export function useLoginController() {
    const { submit } = useLoginHandler();
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    async function onSubmit(input: LoginRequest): Promise<ApiResult<unknown>> {
        setLoading(true);
        setError(null);
        const result = await submit(input);
        setLoading(false);
        if (!result.ok) {
            setError(result.error);
        }
        return result;
    }

    return { onSubmit, loading, error };
}