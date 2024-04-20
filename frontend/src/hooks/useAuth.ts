import { jwtDecode } from "jwt-decode";
import { useEffect, useState } from "react";
import { useAuthStore } from "../storage/auth";

const getLogin = (token: string): [login: string, authorized: boolean] => {
  const payload = jwtDecode<{ login?: string }>(token);
  if (!payload.login) {
    return ["", false];
  }
  return [payload.login, true];
};

export const useAuth = (): { login: string; authorized: boolean } => {
  const [login, setLogin] = useState("");
  const [authorized, setAuthorized] = useState(false);
  const token = useAuthStore((store) => store.token);
  useEffect(() => {
    if (token) {
      const [login, authorized] = getLogin(token);
      setLogin(login);
      setAuthorized(authorized);
    } else {
      setLogin("");
      setAuthorized(false);
    }
  }, [token]);
  return { login, authorized };
};
