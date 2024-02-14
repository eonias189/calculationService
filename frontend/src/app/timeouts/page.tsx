"use client";
import { fetchTimeouts, setTimeouts } from "@/api";
import { contract } from "@/contract";
import { FC, useEffect, useState } from "react";

const Timeouts: FC = () => {
    const [timeouts, SetTimeouts] = useState<contract.Timeouts>();
    useEffect(() => {
        fetchTimeouts().then(SetTimeouts);
    }, []);
    return (
        <>
            <p>timeouts</p>
            <p>Add: {timeouts?.add}</p>
            <p>Substract: {timeouts?.substract}</p>
            <p>Multiply: {timeouts?.multiply}</p>
            <p>Divide: {timeouts?.divide}</p>
            <button
                onClick={() => {
                    const t = new contract.Timeouts();
                    t.add = 4;
                    t.substract = 3;
                    t.multiply = 2;
                    t.divide = 1;
                    setTimeouts(t);
                }}
                className="bg-teal-500 rounded p-1"
            >
                button
            </button>
        </>
    );
};

export default Timeouts;
