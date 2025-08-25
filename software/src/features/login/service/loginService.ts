import { z } from "zod";
import { loginRepository } from "../repository/loginRepository";
import type { LoginRequest, LoginReponse } from "../entitys/loginEntity";
import type { ApiResult } from "../../../app/types";

const schema = z.object({
    login: z.string().min(3, "Login must be more than 3 characters").max(100),
    password: z.string().min(6, "Password must be more than 6 characters").max(100),
});

export const loginService = {
    validate(input: LoginRequest) {
        const result = schema.safeParse(input);
        if(!result.success){
            const msg = result.error.issues[0]?.message ?? "Unknown error";
            return { ok: false as const, error: msg };
        }
        return { ok: true as const};
    },
    async authenticate(input: LoginRequest): Promise<ApiResult<LoginReponse>> {
        const valid = this.validate(input);
        if (!valid.ok) return {ok:false, error: valid.error};

        return loginRepository.login(input);
    }
};

