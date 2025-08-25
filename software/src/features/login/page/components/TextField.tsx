import { InputHTMLAttributes } from "react";

type Props = InputHTMLAttributes<HTMLInputElement> & {
    label: string;
    error?: string | null;
};


export default function TextField({label, error, ...rest}: Props) {
    return (
        <label className="block space-y-1">
            <span className="text-sm text-gray-700">{label}</span>
            <input
                {...rest}
                className={`w-full rounded-lg border px-3 py-2 outline-none focus:ring ${error ? 'border-red-500' : 'border-gray-300'}`}
            />
            {error ? <span className="text-red-500 text-sm">{error}</span> : <span className="text-red-500 text-sm invisible">.</span> /* Placeholder for consistent spacing */}
        </label>
    );
}