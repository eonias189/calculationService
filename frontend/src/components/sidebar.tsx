import { FC } from "react";
import { useAuth } from "../hooks/useAuth";
import { ReactComponent as ProfileIcon } from "../assets/icons/profile.svg";
import { useNavigate } from "react-router-dom";
import Button from "./ui/button";

const Sidebar: FC = () => {
  const { login, authorized, clearToken } = useAuth();
  const navigate = useNavigate();

  return (
    <div className="w-20 flex flex-col items-center pt-[0.6rem]">
      {authorized ? (
        <>
          <ProfileIcon className="size-[3.5rem]" />
          <p className="text-center">{login}</p>
          <Button onClick={() => clearToken()}>log out</Button>
        </>
      ) : (
        <span className="text-sky-500 text-nowrap">
          <Button onClick={() => navigate("/auth")}>log in</Button>
        </span>
      )}
    </div>
  );
};

export default Sidebar;
