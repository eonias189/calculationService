import { FC } from "react";
import { Link, useLocation, useNavigate } from "react-router-dom";
import clsx from "clsx";

interface ILink {
  to: string;
  label: string;
}

const links: ILink[] = [
  {
    to: "/",
    label: "tasks",
  },
  {
    to: "/agents",
    label: "agents",
  },
  {
    to: "/timeouts",
    label: "timeouts",
  },
];

const Navbar: FC = () => {
  const { pathname } = useLocation();

  return (
    <div className="flex flex-row gap-4 w-[100%] p-[0.8rem]">
      {links.map((link) => (
        <span
          key={link.to}
          className={clsx("text-sky-500 border-sky-500 px-[0.4rem] h-min rounded hover:border-b-4 transition", {
            "border-b-4": pathname === link.to,
          })}
        >
          <Link to={link.to}>{link.label}</Link>
        </span>
      ))}
    </div>
  );
};

export default Navbar;
