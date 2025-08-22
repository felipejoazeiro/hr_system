import LoginForm from "./components/LoginForm";

export default function LoginPage() {
    return (
        <div className="grid w-screen place-items-center">
            <div className="max-w-md rounded-2xl bg-white p-6 shadow-lg w-64">
                <h1 className="mb-4 text-xl font-semibold text-black">System Access</h1>
                <LoginForm />
            </div>
        </div>
    );
}
