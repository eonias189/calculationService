"use client";
import { addTask, fetchTasks } from "@/api";
import Task from "@/components/task";
import Button from "@/components/ui/button";
import Input from "@/components/ui/input";
import { FC, useState } from "react";
import { useFetcher } from "@/hooks/useFetcher";
import Loader from "@/components/loader";

const Index: FC = () => {
    const tasksFetcher = useFetcher(fetchTasks);
    const [expression, setExpression] = useState("");
    return (
        <>
            <h1 className="text-center font-normal text-2xl mt-4">Tasks</h1>
            <div className="flex flex-col gap-1 ml-2">
                <div className="flex flex-row gap-2 my-3">
                    <Input
                        value={expression}
                        placeholder="expression"
                        onChange={(e) => {
                            e.preventDefault();
                            setExpression(e.currentTarget.value);
                        }}
                    />
                    <Button
                        className="w-32"
                        onClick={() => {
                            addTask(expression)
                                .then(() => setExpression(""))
                                .then(() => {
                                    setTimeout(() => tasksFetcher.update(), 500);
                                })
                                .then(() => {
                                    setTimeout(() => tasksFetcher.update(), 2500);
                                });
                        }}
                    >
                        Add
                    </Button>
                </div>
                {tasksFetcher.isLoading ? (
                    <Loader className="mx-auto mt-20" />
                ) : (
                    tasksFetcher.data?.map((task) => <Task key={task.id} task={task} />)
                )}
            </div>
        </>
    );
};

export default Index;
