import { useState } from "react";
import TextField from "./TextField";
import PasswordField from "./PasswordField";
import ErrorBanner from "./ErrorBanner";
import { useLoginController } from "../../application/loginController"; 
import type { LoginRequest } from "../../entitys/loginEntity";
import { set } from "zod";

export default function LoginForm() {
    const { onSubmit, loading, error} = useLoginController();
    const [values, setValues] = useState<LoginRequest>({
        login: "",
        password: ""
    });

    const [fieldErr, setFieldErr] = useState<Partial<Record<keyof LoginRequest, string>>>({});

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        setValues((v)=> ({...v, [e.target.name]: e.target.value}));
    }

    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        const errs: Partial<Record<keyof LoginRequest, string>> = {};
        if (!values.login) errs.login = "Login is required";
        if (!values.password) errs.password = "Password is required";

        setFieldErr(errs);
        if (Object.keys(errs).length) return;
        await onSubmit(values);
    }

    return (
        <form onSubmit={handleSubmit} className="space-y-3">
            {error && <ErrorBanner error={error} />}
            <TextField
                name="login"
                label="Login"
                value={values.login}
                onChange={handleChange}
                error={fieldErr.login ?? null}
            />
            <PasswordField
                name="password"
                label="Password"
                value={values.password}
                onChange={handleChange}
                error={fieldErr.password ?? null}
            />
            <button type="submit" disabled={loading} className="mt-2 w-full rounded-lg bg-black px-4 py-2 text-white disabled:opacity-60">
                {loading ? "Logging in..." : "Log in"}
            </button>
        </form>
    );

}