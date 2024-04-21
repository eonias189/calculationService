import { FC, useState } from "react";
import Input from "../ui/input";
import Button from "../ui/button";
import Form from "../ui/form";

export type LoginFormFields = {
  login: string;
  password: string;
};

export interface LoginFormProps {
  onSubmit(data: LoginFormFields): Promise<void>;
}

const LoginForm: FC<LoginFormProps> = ({ onSubmit }) => {
  const [data, setData] = useState<LoginFormFields>({ login: "", password: "" });

  const onFieldChange = (field: keyof LoginFormFields): React.ChangeEventHandler<HTMLInputElement> => {
    return (e) => {
      e.preventDefault();
      setData((last) => ({ ...last, [field]: e.target.value }));
    };
  };

  return (
    <Form>
      <Input
        type="text"
        name="login"
        required
        value={data.login}
        placeholder="login"
        onChange={onFieldChange("login")}
      />
      <Input
        type="password"
        name="password"
        required
        value={data.password}
        placeholder="password"
        onChange={onFieldChange("password")}
      />
      <Button
        onClick={async (e) => {
          e.preventDefault();
          await onSubmit(data);
        }}
      >
        log in
      </Button>
    </Form>
  );
};

export default LoginForm;
