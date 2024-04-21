import { create } from "zustand";

export type AuthStore = {
  token?: string;
  setToken(token: string): void;
  clearToken(): void;
};

export const useAuthStore = create<AuthStore>((set) => {
  return {
    setToken: (token: string): void => {
      set((store) => ({ token }));
    },
    clearToken: (): void => {
      set((store) => ({ token: undefined }));
    },
  };
});
