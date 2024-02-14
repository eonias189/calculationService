"use client";
import { addTask, fetchTasks } from "@/api";
import { contract } from "@/contract";
import { FC, useEffect, useState } from "react";

const Index: FC = () => {
    const [tasks, setTasks] = useState<contract.Task[]>();
    useEffect(() => {
        fetchTasks().then(setTasks);
    }, []);
    return (
        <>
            <p>tasks</p>
            {tasks?.map((task) => (
                <p key={task.id} className="">
                    {task.expression}
                </p>
            ))}
            <button onClick={() => addTask("2 + 2 * 2")} className="bg-teal-500 rounded p-1">
                btn
            </button>
        </>
    );
};

export default Index;
