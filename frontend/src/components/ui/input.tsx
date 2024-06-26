import clsx from "clsx";
import { FC, InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {}

const Input: FC<InputProps> = ({ className, ...props }) => {
  return (
    <input
      className={clsx(`w-[80%] border-solid border-primary border-2 rounded-lg px-[0.6rem] py-[0.3rem] outline-none`, {
        [className!]: className !== undefined,
      })}
      {...props}
    />
  );
};

export default Input;
