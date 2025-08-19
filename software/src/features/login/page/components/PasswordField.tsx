import { useState } from 'react';
import TextField from './TextField';

type Props = {
    label: string;
    name?: string;
    value?: string;
    onChange: React.ChangeEventHandler<HTMLInputElement>;
    error?: string | null;
}

export default function PasswordField(props: Props) {
    const [show, setShow] = useState(false);
    return (
        <div className = "relative">
            <TextField {...props} type={show ? "text" : "password"}/>
            <button type='button' onClick={() => setShow((s) => !s)} className='absolute right-2 top-1/2 -translate-y-1/2'>
                {show ? "Hide" : "Show"}
            </button>
        </div>
    );
}