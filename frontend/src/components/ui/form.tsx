import React, { FC, useState } from "react";
import Input from "./input";

export interface FormField {
  placeholder?: string;
  formType: React.HTMLInputTypeAttribute;
}

export interface FormProps {
  fields: { [name: string]: FormField };
}

const Form: FC<FormProps> = ({ fields }) => {
  const [data, setData] = useState<typeof fields>();
  return <form></form>;
};
