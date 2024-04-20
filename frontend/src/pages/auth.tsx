import { FC, useEffect, useState, HTMLInputTypeAttribute } from "react";
import { login } from "../api";
import { useAuthStore } from "../storage/auth";
import { useAuth } from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";
import Input from "../components/ui/input";

const AuthPage: FC = () => {
  const [authType, setAuthType] = useState<"login" | "register">("login");
  const { setToken, clearToken } = useAuthStore();
  const { authorized } = useAuth();
  const navigate = useNavigate();
  useEffect(() => {
    console.log(authorized);
    if (authorized) {
      navigate("/");
    }
  }, [authorized]);

  return (
    <div>
      <h1>log in page</h1>
      <button
        onClick={() => {
          login("lala", "lolo")
            .then((token) => {
              setToken(token);
            })
            .catch(alert);
        }}
      >
        log in
      </button>
      <button onClick={() => clearToken()}>log out</button>
    </div>
  );
};

export default AuthPage;
