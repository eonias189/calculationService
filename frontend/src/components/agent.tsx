import { contract } from "@/contract";
import { FC } from "react";
import cx from "classnames";

interface AgentProps {
    agent: contract.AgentData;
}

const Agent: FC<AgentProps> = ({ agent }) => {
    return (
        <div className="border-2 bg-slate-300 p-2 text-nowrap w-2/3 rounded-xl flex flex-row items-center">
            <span
                className={`size-4 rounded-full mr-2 ${cx({
                    "bg-lime-500": agent.ping < 100,
                    "bg-amber-500": agent.ping >= 100 && agent.ping < 800,
                    "bg-red-500": agent.ping >= 800,
                })}`}
            ></span>
            <p>
                ping: {agent.ping}; threads: {agent.status.executingThreads} of {agent.status.maxThreads}
            </p>
        </div>
    );
};

export default Agent;
