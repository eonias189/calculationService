import { FC, ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {}

const Button: FC<ButtonProps> = ({ children, className, ...props }) => {
    return (
        <button className={`bg-teal-500 rounded-md p-1 ${className}`} {...props}>
            {children}
        </button>
    );
};

export default Button;
