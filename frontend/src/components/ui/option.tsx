import clsx from "clsx";
import { FC, ReactNode } from "react";

export interface OptionProps {
  children?: ReactNode;
  active: boolean;
  onClick?: React.MouseEventHandler<HTMLSpanElement>;
  className?: string;
}

const Option: FC<OptionProps> = ({ onClick, children, active, className }) => {
  return (
    <span
      className={clsx(
        "text-primary border-primary px-[0.4rem] h-min hover:border-b-[3px] transition hover:cursor-pointer",
        {
          "border-b-[3px]": active,
          [className!]: className !== undefined,
        }
      )}
      onClick={onClick}
    >
      {children}
    </span>
  );
};

export default Option;
