import { FC, useEffect, useState } from "react";
import { login, register } from "../api";
import { useAuth } from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";
import LoginForm, { LoginFormFields } from "../components/forms/loginForm";
import Option from "../components/ui/option";
import RegisterForm, { RegisterFormFields } from "../components/forms/registerForm";

const AuthPage: FC = () => {
  const [authType, setAuthType] = useState<"login" | "register">("login");
  const [error, setError] = useState("");
  const { authorized, setToken } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (authorized) {
      navigate("/");
    }
  }, [authorized, navigate]);

  const onLogin = async (data: LoginFormFields): Promise<void> => {
    setError("");
    return login(data.login, data.password)
      .then((token) => setToken(token))
      .then(() => navigate("/"))
      .catch((err) => setError(`${err}`));
  };

  const onRegister = async (data: RegisterFormFields): Promise<void> => {
    setError("");
    if (data.password !== data.repeatPassword) {
      setError("Passwords must coincide");
      return;
    }
    return register(data.login, data.password)
      .then(() => onLogin(data))
      .catch((err) => setError(`${err}`));
  };

  return (
    <div className="w-[100%] h-[100%] flex flex-col items-center pt-[5rem]">
      <div className="w-full max-w-[50%] min-w-[20rem] min-h-[60%] flex flex-col gap-[1.6rem] border-primary border-2 rounded-lg">
        <div className="w-[100%] flex flex-row">
          <Option active={authType === "login"} className="w-[50%] text-center" onClick={() => setAuthType("login")}>
            log in
          </Option>
          <Option
            active={authType === "register"}
            className="w-[50%] text-center"
            onClick={() => setAuthType("register")}
          >
            register
          </Option>
        </div>
        {authType === "login" ? (
          <LoginForm onSubmit={onLogin}></LoginForm>
        ) : (
          <RegisterForm onSubmit={onRegister}></RegisterForm>
        )}
        <div className="p-[0.5rem] border-primary border-t-2 text-sm">
          <p>{error}</p>
        </div>
      </div>
    </div>
  );
};

export default AuthPage;
