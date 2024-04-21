import clsx from "clsx";
import { FC, HTMLAttributes } from "react";

export interface HeaderProps extends HTMLAttributes<HTMLParagraphElement> {}

const Header: FC<HeaderProps> = ({ className, children, ...props }) => {
  return (
    <h1
      className={clsx("text-center w-[100%] text-xl mt-[1rem]", {
        [className!]: className !== undefined,
      })}
    >
      {children}
    </h1>
  );
};

export default Header;
