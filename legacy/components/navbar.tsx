"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import cx from "classnames";
import { FC } from "react";

type LinkData = {
    name: string;
    href: string;
};

interface NavbarProps {
    links: LinkData[];
}

const Navbar: FC<NavbarProps> = ({ links }) => {
    const pathName = usePathname();
    return (
        <nav className="flex flex-row mx-auto mt-2 gap-2  w-4/5 rounded-lg">
            {links.map((link, key) => (
                <Link
                    key={key}
                    href={link.href}
                    className={`rounded-md hover:border-teal-500 p-1 relative overflow-hidden after:size-full after:absolute after:top-0 after:left-0 after:-translate-x-full after:z-10 after:bg-teal-500  hover:after:translate-x-0 after:transition-all after:duration-300 ${cx(
                        {
                            "bg-teal-500": pathName === link.href,
                        }
                    )}`}
                >
                    <span className="relative z-20 -translate-x-1/2 -translate-y-1/2">{link.name}</span>
                </Link>
            ))}
        </nav>
    );
};

export default Navbar;
