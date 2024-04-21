import { FC, useState } from "react";
import { Timeouts } from "../../api/types";
import Input from "./input";
import Button from "./button";

export interface TimeoutsFormProps {
  data: Timeouts;
  onSave: (timeouts: Timeouts) => Promise<void>;
}

function Keys<T extends Record<any, any>>(obj: T): (keyof T)[] {
  const res: (keyof T)[] = [];
  for (let key in obj) {
    res.push(key);
  }
  return res;
}

const timeoutsBabel: { [x in keyof Timeouts]: string } = {
  add: "+",
  sub: "-",
  mul: "*",
  div: "/",
};

const TimeoutsUI: FC<TimeoutsFormProps> = ({ data, onSave }) => {
  const [timeouts, setTimeouts] = useState(data);

  const onTimeoutChange = (field: keyof Timeouts): React.ChangeEventHandler<HTMLInputElement> => {
    return (e) => {
      e.preventDefault();
      const value = parseInt(e.target.value);
      setTimeouts((last) => ({ ...last, [field]: value }));
    };
  };

  return (
    <form className="min-w-[16rem] max-w-[70%] w-full flex flex-col pl-[1rem]">
      {Keys(timeouts).map((field) => (
        <div key={field} className="flex flex-row gap-[0.4rem]">
          <p>{timeoutsBabel[field]}</p>
          <Input name={field} type="range" onChange={onTimeoutChange(field)} value={timeouts[field]} />
          <p>{timeouts[field]}</p>
        </div>
      ))}
      <Button
        className="w-min self-center"
        onClick={(e) => {
          e.preventDefault();
          onSave(timeouts);
        }}
      >
        save
      </Button>
    </form>
  );
};

export default TimeoutsUI;
