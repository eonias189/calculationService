import { FC, useState } from "react";
import { useQuery } from "react-query";
import { getAgents } from "../api";
import AgentUi from "../components/ui/agent";
import Input from "../components/ui/input";
import Header from "../components/ui/header";

const AgentsPage: FC = () => {
  const [showDisable, setShowDisable] = useState(false);

  const {
    data: agents,
    isFetching,
    isError,
    error,
  } = useQuery("agents", async () => getAgents({ limit: 99999 }), {
    cacheTime: 0,
    retry: false,
  });

  if (isError) {
    return <Header>error while fetching agents: {`${error}`}</Header>;
  }

  if (isFetching) {
    return <Header>Fetching</Header>;
  }

  return (
    <div className="w-[100%] h-[100%]] flex flex-col items-start ml-[1.5rem] gap-[0.6rem]">
      <div className="flex flex-row gap-[0.3rem] items-center">
        <Input
          name="show-disable"
          type="checkbox"
          checked={showDisable}
          className="w-[1rem] h-[1rem]"
          onChange={() => setShowDisable((last) => !last)}
        />
        <label htmlFor="show-disable" className="text-nowrap">
          show disable
        </label>
      </div>

      {agents
        ?.filter((agent) => showDisable || agent.active)
        ?.map((agent) => (
          <AgentUi key={agent.id} agent={agent} />
        ))}
    </div>
  );
};

export default AgentsPage;
