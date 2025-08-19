export default function ErrorBanner({ error }: { error: string | null }) {
    if (!error) return null;

    return (
        <div className="rounded-lg border border-red-300 bg-red-50 px-3 py-2 text-sm text-red-700">
            {error}
        </div>
    );
}
