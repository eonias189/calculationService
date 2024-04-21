import { FC, useState } from "react";
import Input from "../ui/input";
import Button from "../ui/button";
import Form from "../ui/form";

export type RegisterFormFields = {
  login: string;
  password: string;
  repeatPassword: string;
};

export interface RegisterFormProps {
  onSubmit(data: RegisterFormFields): Promise<void>;
}

const RegisterForm: FC<RegisterFormProps> = ({ onSubmit }) => {
  const [data, setData] = useState<RegisterFormFields>({ login: "", password: "", repeatPassword: "" });

  const onFieldChange = (field: keyof RegisterFormFields): React.ChangeEventHandler<HTMLInputElement> => {
    return (e) => {
      e.preventDefault();
      setData((last) => ({ ...last, [field]: e.target.value }));
    };
  };

  return (
    <Form>
      <Input type="text" name="login" placeholder="login" value={data.login} onChange={onFieldChange("login")} />
      <Input
        type="password"
        name="password"
        placeholder="password"
        value={data.password}
        onChange={onFieldChange("password")}
      />
      <Input
        name="repeate password"
        type="password"
        placeholder="repeat password"
        value={data.repeatPassword}
        onChange={onFieldChange("repeatPassword")}
      />
      <Button
        onClick={async (e) => {
          e.preventDefault();
          await onSubmit(data);
        }}
      >
        register
      </Button>
    </Form>
  );
};

export default RegisterForm;
