import { post } from "../../../app/http/httpClient";
import type { LoginRequest, LoginReponse } from "../entitys/loginEntity";
import type { ApiResult } from "../../../app/types";

export const loginRepository = {
    async login(payload: LoginRequest): Promise<ApiResult<LoginReponse>> {
        try{
            const data = await post<LoginReponse>("/login", payload);
            return { ok: true, data};
        }catch (e : any) {
            return { ok: false, error: e.message ?? "Unknown error" };
        }
    }
}