import { FC, ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {}

const Button: FC<ButtonProps> = ({ children, className, ...props }) => {
  return (
    <button className={`bg-primary text-white font-normal rounded-md px-2 py-1 text-nowrap ${className}`} {...props}>
      {children}
    </button>
  );
};

export default Button;
