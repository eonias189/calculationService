"use client";
import { fetchTimeouts, setTimeouts as sendTimeouts } from "@/api";
import { contract } from "@/contract";
import { useFetcher } from "@/hooks/useFetcher";
import Input from "@/components/ui/input";
import { FC, useEffect, useState } from "react";
import Button from "@/components/ui/button";
import Loader from "@/components/loader";

const entries = <T extends Object>(obj: T): [keyof T, T[keyof T]][] => {
    const res: [keyof T, T[keyof T]][] = [];
    for (let key in obj) {
        res.push([key, obj[key]]);
    }
    return res;
};

const Timeouts: FC = () => {
    const timeoutsFetcher = useFetcher(fetchTimeouts);
    const [timeouts, setTimeouts] = useState(new contract.Timeouts().toObject());
    const names: Record<keyof typeof timeouts, string> = {
        add: "+",
        substract: "-",
        multiply: "*",
        divide: "/",
    };
    useEffect(() => {
        setTimeouts({
            ...timeouts,
            add: timeoutsFetcher.data?.add ?? 0,
            substract: timeoutsFetcher.data?.substract ?? 0,
            multiply: timeoutsFetcher.data?.multiply ?? 0,
            divide: timeoutsFetcher.data?.divide ?? 0,
        });
    }, [timeoutsFetcher.data]);
    return (
        <>
            <h1 className="text-center font-normal text-2xl mt-4">Timeouts</h1>
            <form className="flex flex-col gap-1 m-2">
                {timeoutsFetcher.isLoading ? (
                    <Loader className="mx-auto mt-32" />
                ) : (
                    <>
                        {entries(timeouts).map(([key, value], num) => (
                            <div key={num} className="flex flex-row items-center gap-2">
                                <label>{names[key]}</label>
                                <Input
                                    type="range"
                                    max={100}
                                    value={value}
                                    onChange={(e) => {
                                        e.preventDefault();
                                        setTimeouts({
                                            ...timeouts,
                                            [key]: parseInt(e.currentTarget.value || "0"),
                                        });
                                    }}
                                />
                                <label>{value}</label>
                            </div>
                        ))}
                        <Button
                            onClick={(e) => {
                                e.preventDefault();
                                sendTimeouts(contract.Timeouts.fromObject(timeouts)).then(() => {
                                    setTimeout(() => timeoutsFetcher.update(), 500);
                                });
                            }}
                            className="w-min text-nowrap"
                        >
                            set timeouts
                        </Button>
                    </>
                )}
            </form>
        </>
    );
};

export default Timeouts;
