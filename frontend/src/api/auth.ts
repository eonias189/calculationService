import { BASE_URL } from ".";
import { RegisterReq, LoginReq, LoginResp, ErrorResp } from "./types";

export const register = async (login: string, password: string): Promise<void> => {
  const url = BASE_URL + "/auth/register";
  const body: RegisterReq = { login, password };
  const resp = await fetch(url, { method: "POST", body: JSON.stringify(body) });
  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }
};

export const login = async (login: string, password: string): Promise<string> => {
  const url = BASE_URL + "/auth/login";
  const body: LoginReq = { login, password };
  const resp = await fetch(url, { method: "POST", body: JSON.stringify(body) });
  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }
  const res = (await resp.json()) as LoginResp;
  return res.token;
};
