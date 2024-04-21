import { FC, useEffect, useState } from "react";
import { Task } from "../api/types";
import { addTask, getTasks } from "../api";
import { useAuth } from "../hooks/useAuth";
import { useMutation, useQuery, useQueryClient } from "react-query";
import TaskUI from "../components/ui/task";
import Input from "../components/ui/input";
import Button from "../components/ui/button";

const TasksPage: FC = () => {
  const [expression, setExpression] = useState("");
  const { token } = useAuth();
  const queryClient = useQueryClient();

  const {
    data: tasks,
    isError,
    isFetching,
  } = useQuery(
    ["tasks", token],
    async () => {
      return await getTasks(token ?? "", { limit: 999999 });
    },
    { retry: false }
  );

  const onAddTask = useMutation(() => addTask(expression, token ?? ""), {
    onSuccess: () => {
      queryClient.invalidateQueries("tasks");
    },
  });

  if (isError) {
    return <h1>Необходима авторизация</h1>;
  }

  return (
    <div className="w-[100%] h-[100%]] flex flex-col items-start gap-[1.4rem]">
      <div className="w-[50%] ml-[1.5rem] mt-[0.6rem] flex flex-row gap-[0.5rem]">
        <Input
          name="new_task"
          placeholder="new task"
          value={expression}
          onChange={(e) => {
            e.preventDefault();
            setExpression(e.target.value);
          }}
        />
        <Button onClick={() => onAddTask.mutate()}>add</Button>
      </div>
      <div className="ml-[1.5rem] flex flex-col gap-[0.5rem]">
        {tasks?.map((task) => (
          <TaskUI key={task.id} task={task} />
        ))}
      </div>
    </div>
  );
};

export default TasksPage;
