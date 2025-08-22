import { useState } from "react";
import { useNavigate } from "react-router-dom";
import TextField from "./TextField";
import PasswordField from "./PasswordField";
import ErrorBanner from "./ErrorBanner";
import { useLoginController } from "../../application/loginController"; 
import type { LoginRequest } from "../../entitys/loginEntity";
import { set } from "zod";

export default function LoginForm() {
    const { onSubmit, loading, error, blockSeconds } = useLoginController();
    const navigate = useNavigate();
    const [values, setValues] = useState<LoginRequest>({
        login: "",
        password: ""
    });

    const [fieldErr, setFieldErr] = useState<Partial<Record<keyof LoginRequest, string>>>({});

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target;
        setValues((v) => ({ ...v, [name]: value }));
        // Limpa o erro do campo ao digitar
        setFieldErr((prev) => ({ ...prev, [name]: undefined }));
    }

    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        const errs: Partial<Record<keyof LoginRequest, string>> = {};
        if (!values.login) errs.login = "Login is required";
        if (!values.password) errs.password = "Password is required";

        setFieldErr(errs);
        if (Object.keys(errs).length) return;
        const result = await onSubmit(values);
        if (result?.ok) {
            // Redireciona para a página principal após login
            navigate("/home"); 
        }
    }

    return (
        <form onSubmit={handleSubmit} className="space-y-3">
            {error && (
                error.includes("Too many failed attempts") ? (
                    <ErrorBanner error={`Too many failed attempts. Please try again in ${blockSeconds} seconds.`} />
                ) : (
                    <ErrorBanner error={error} />
                )
            )}
            <TextField
                name="login"
                label="Login"
                value={values.login}
                onChange={handleChange}
                error={fieldErr.login ?? null}
                autoFocus
            />
            <PasswordField
                name="password"
                label="Password"
                value={values.password}
                onChange={handleChange}
                error={fieldErr.password ?? null}
            />
            <button type="submit" disabled={loading} className="mt-2 w-full rounded-lg bg-black px-4 py-2 text-white disabled:opacity-60">
                {loading ? <> <span className="inline-block w-4 h-4 mr-2 border-2 border-white border-t-transparent rounded-full animate-spin"></span> </> : "Log in"}
            </button>
        </form>
    );

}