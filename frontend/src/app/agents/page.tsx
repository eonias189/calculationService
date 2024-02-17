"use client";
import { fetchAgents } from "@/api";
import { contract } from "@/contract";
import { FC, useState, useEffect } from "react";

const Agents: FC = () => {
    const [agents, setAgents] = useState<contract.AgentData[]>();
    useEffect(() => {
        fetchAgents().then(setAgents);
    }, []);
    return (
        <>
            <p>agents</p>
            {agents?.map((agent) => (
                <p key={agent.id}>
                    {agent.ping}, {agent.status.executingThreads}, {agent.status.maxThreads}
                </p>
            ))}
        </>
    );
};

export default Agents;
