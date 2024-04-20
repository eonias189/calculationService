import { FC, InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {}

const Input: FC<InputProps> = ({ className, ...props }) => {
    return (
        <input
            className={`border-solid border-teal-600 border-2 rounded-lg p-1 outline-none w-8/12`}
            {...props}
        />
    );
};

export default Input;
