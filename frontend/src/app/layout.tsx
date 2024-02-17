import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Navbar from "@/components/navbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Calculation Service",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body className={inter.className}>
                <Navbar
                    links={[
                        { name: "tasks", href: "/" },
                        { name: "servers", href: "/agents" },
                        { name: "timeouts", href: "/timeouts" },
                    ]}
                />
                {children}
            </body>
        </html>
    );
}
