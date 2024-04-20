import { FC } from "react";
import { useAuth } from "../hooks/useAuth";
import { useAuthStore } from "../storage/auth";
import { ReactComponent as ProfileIcon } from "../assets/icons/profile.svg";
import { Link } from "react-router-dom";
import Button from "./ui/button";

const Sidebar: FC = () => {
  const { login, authorized } = useAuth();
  const { clearToken } = useAuthStore();

  return (
    <div className="px-4 pt-2">
      {authorized ? (
        <div className="flex flex-col content-center">
          <ProfileIcon className="size-[3.5rem]" />
          <p className="text-center">{login}</p>
          <Button onClick={() => clearToken()}>log out</Button>
        </div>
      ) : (
        <span className="text-sky-500 text-nowrap">
          <Link to="/auth">log in</Link>
        </span>
      )}
    </div>
  );
};

export default Sidebar;
