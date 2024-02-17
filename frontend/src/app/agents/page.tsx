"use client";
import { fetchAgents } from "@/api";
import Agent from "@/components/agent";
import { contract } from "@/contract";
import { useFetcher } from "@/hooks/useFetcher";
import { FC, useState, useEffect } from "react";

const Agents: FC = () => {
    const agentsFetcher = useFetcher(fetchAgents);
    useEffect(() => {
        fetchAgents().then(agentsFetcher.update);
    }, []);
    return (
        <>
            <h1 className="text-center font-normal text-2xl mt-4">Servers</h1>
            <div className="flex flex-col gap-1 ml-2 items-center">
                {agentsFetcher.data?.map((agent) => (
                    <Agent key={agent.id} agent={agent} />
                ))}
            </div>
        </>
    );
};

export default Agents;
