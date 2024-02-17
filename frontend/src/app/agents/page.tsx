"use client";
import { fetchAgents } from "@/api";
import Agent from "@/components/agent";
import Loader from "@/components/loader";
import { useFetcher } from "@/hooks/useFetcher";
import { FC } from "react";

const Agents: FC = () => {
    const agentsFetcher = useFetcher(fetchAgents);
    return (
        <>
            <h1 className="text-center font-normal text-2xl mt-4">Servers</h1>
            <div className="flex flex-col gap-1 mt-8 items-center">
                {agentsFetcher.isLoading ? (
                    <Loader className="mx-auto mt-32" />
                ) : (
                    agentsFetcher.data?.map((agent) => <Agent key={agent.id} agent={agent} />)
                )}
            </div>
        </>
    );
};

export default Agents;
