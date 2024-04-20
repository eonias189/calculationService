import { create } from "zustand";

export type AuthStore = {
  token?: string;
  setToken(token: string): void;
  clearToken(): void;
};

export const useAuthStore = create<AuthStore>((set) => {
  const tokenWas = localStorage.getItem("token");
  return {
    token: tokenWas === null ? undefined : tokenWas,
    setToken: (token: string): void => {
      localStorage.setItem("token", token);
      set((store) => ({ token }));
    },
    clearToken: (): void => {
      localStorage.removeItem("token");
      set((store) => ({ token: undefined }));
    },
  };
});
