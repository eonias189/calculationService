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
        <nav className="flex flex-row">
            {links.map((link, key) => (
                <Link
                    key={key}
                    href={link.href}
                    className={`border-solid border-2 border-teal-600 rounded-xl mx-2 mt-2 px-1 ${cx({
                        "bg-teal-600": pathName === link.href,
                    })}`}
                >
                    {link.name}
                </Link>
            ))}
        </nav>
    );
};

export default Navbar;
