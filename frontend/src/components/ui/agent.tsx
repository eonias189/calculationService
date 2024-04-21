import { FC } from "react";
import { Agent } from "../../api/types";
import clsx from "clsx";

interface AgentUiProps {
  agent: Agent;
}

const AgentUi: FC<AgentUiProps> = ({ agent }) => {
  return (
    <div className="border-2 bg-slate-300 p-2 text-nowrap w-2/3 rounded-xl flex flex-row items-center">
      <span
        className={`size-4 rounded-full mr-2 ${clsx({
          "bg-lime-500": agent.active,
          "bg-red-500": !agent.active,
        })}`}
      ></span>
      <p>
        id: {agent.id}; ping: {agent.ping}ms; running {agent.runningThreads} threads of {agent.maxThreads}
      </p>
    </div>
  );
};

export default AgentUi;
