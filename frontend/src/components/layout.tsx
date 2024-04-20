import { FC, ReactElement, ReactNode } from "react";

export interface LayoutProps {
  navbar: ReactNode;
  sidebar: ReactNode;
  children?: ReactNode;
}

const Layout: FC<LayoutProps> = ({ navbar, sidebar, children }) => {
  return (
    <div className="w-[100vw] h-[100vh] grid grid-cols-custom-layout grid-rows-cutom-layout">
      <nav className="flex row-start-1 row-span-1 col-start-1 col-span-1">{navbar}</nav>
      <aside className="flex row-start-1 row-span-2 col-start-2 col-span-1">{sidebar}</aside>
      <main className="flex row-start-2 row-span-1 col-start-1 col-span-1">{children}</main>
    </div>
  );
};

export default Layout;
