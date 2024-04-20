import { contract } from "@/contract";
import { FC } from "react";
import cx from "classnames";

interface TaskProps {
    task: contract.Task;
}

const Task: FC<TaskProps> = ({ task }) => {
    const getResult = (): string => {
        switch (task.status) {
            case contract.TaskStatus.done:
                return task.result.toString();
            case contract.TaskStatus.execution:
            case contract.TaskStatus.pending:
                return "?";
            case contract.TaskStatus.executionError:
                return "NaN";
        }
    };
    return (
        <div
            className={`text-nowrap w-min p-1 border rounded-md ${cx({
                "bg-zinc-600": task.status === contract.TaskStatus.pending,
                "bg-amber-400": task.status === contract.TaskStatus.execution,
                "bg-lime-500": task.status === contract.TaskStatus.done,
                "bg-red-600": task.status === contract.TaskStatus.executionError,
            })}`}
        >
            {task.expression}={getResult()}
        </div>
    );
};

export default Task;
