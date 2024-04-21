import { FC } from "react";
import { Task, TaskStatus } from "../../api/types";
import clsx from "clsx";

export interface TaskProps {
  task: Task;
}

const getResult = (task: Task): string => {
  switch (task.status) {
    case TaskStatus.success:
      return task.result.toString();
    case TaskStatus.executing:
    case TaskStatus.pending:
      return "?";
    case TaskStatus.error:
      return task.result === 0 ? "NaN" : "very big";
  }
};

const TaskUI: FC<TaskProps> = ({ task }) => {
  return (
    <div
      className={`text-nowrap w-min p-1 border rounded-md ${clsx({
        "bg-zinc-600": task.status === TaskStatus.pending,
        "bg-amber-400": task.status === TaskStatus.executing,
        "bg-lime-500": task.status === TaskStatus.success,
        "bg-red-600": task.status === TaskStatus.error,
      })}`}
    >
      {task.expression}={getResult(task)}
    </div>
  );
};

export default TaskUI;
