import { post } from "../../../app/http/httpClient";
import type { LoginRequest, LoginReponse } from "../entitys/loginEntity";
import type { ApiResult } from "../../../app/types";

export const loginRepository = {
    async login(payload: LoginRequest): Promise<ApiResult<LoginReponse>> {
        try{
            //console.log("LoginRepository: Attempting to login with payload:", payload);
            const data = await post<LoginReponse>("/login", payload);
            return { ok: true, data};
        }catch (e : any) {
            //console.error("LoginRepository: Error logging in:", e);
            let errorMsg = "Unknown error";
            if(e.message?.includes("Failed to fetch")) {
                errorMsg = "Network error: Failed to connect to the server";
            }
            return { ok: false, error: errorMsg };
        }
    }
}