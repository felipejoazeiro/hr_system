import LoginForm from "./components/LoginForm";

export default function LoginPage() {
    return (
        <div className="grid min-h-screen place-items-center bg-gray-50">
            <div className="w-full max-w-md rounded-2xl bg-white p-6 shadow-lg">
                <h1 className="mb-4 text-xl font-semibold">System Access</h1>
                <LoginForm />
            </div>
        </div>
    );
}
