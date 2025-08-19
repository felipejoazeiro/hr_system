import { loginService } from "../service/loginService";
import { useAuthStore } from "../../../app/store/authStore";
import type { LoginRequest } from "../entitys/loginEntity";
import type { ApiResult } from "../../../app/types";    


export function useLoginHandler() {
    const setSession = useAuthStore((state)=> state.setSession);

    async function submit(input: LoginRequest): Promise<ApiResult<unknown>> {
        const result = await loginService.authenticate(input);
        if (!result.ok) return {ok: false, error: result.error};
        const { token, position, expires_in } = result.data;
        setSession( token, position, expires_in );
        return { ok: true, data: null };
    }
    return { submit };
}