import clsx from "clsx";
import { FC, FormHTMLAttributes } from "react";

export interface FormProps extends FormHTMLAttributes<HTMLFormElement> {}

const Form: FC<FormProps> = ({ children, className, ...props }) => {
  return (
    <form
      className={clsx("flex flex-col w-[100%] h-[100%] p-[0.5rem] gap-[1rem] items-center", {
        [className!]: className !== undefined,
      })}
      {...props}
    >
      {children}
    </form>
  );
};

export default Form;
