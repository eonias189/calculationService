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

export const LOCAL_STORAGE_TOKEN_KEY = "token";

export const useAuth = (): {
  token?: string;
  login: string;
  authorized: boolean;
  setToken(token: string): void;
  clearToken(): void;
} => {
  const [login, setLogin] = useState("");
  const [authorized, setAuthorized] = useState(false);
  const { token, setToken: set, clearToken: clear } = useAuthStore();

  useEffect(() => {
    const token = localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY);
    if (token) {
      set(token);
    }
  }, [set]);

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

  const setToken = (token: string): void => {
    localStorage.setItem(LOCAL_STORAGE_TOKEN_KEY, token);
    set(token);
  };

  const clearToken = (): void => {
    localStorage.removeItem(LOCAL_STORAGE_TOKEN_KEY);
    clear();
  };

  return { token, login, authorized, setToken, clearToken };
};
