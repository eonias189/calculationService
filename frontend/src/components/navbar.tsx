import { FC } from "react";
import { Link, useLocation } from "react-router-dom";
import Option from "./ui/option";

export interface ILink {
  to: string;
  label: string;
}

export interface NavbarProps {
  links: ILink[];
}

const Navbar: FC<NavbarProps> = ({ links }) => {
  const { pathname } = useLocation();

  return (
    <div className="flex flex-row gap-4 w-[100%] p-[0.8rem]">
      {links.map((link) => (
        <Option key={link.to} active={pathname === link.to}>
          <Link to={link.to}>{link.label}</Link>
        </Option>
      ))}
    </div>
  );
};

export default Navbar;
