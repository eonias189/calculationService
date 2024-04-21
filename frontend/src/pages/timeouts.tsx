import { FC } from "react";
import { useQuery, useQueryClient } from "react-query";
import { getTimeouts, setTimeouts } from "../api";
import { useAuth } from "../hooks/useAuth";
import Header from "../components/ui/header";
import TimeoutsUI from "../components/ui/timeouts";
import { Timeouts } from "../api/types";

const TimeoutsPage: FC = () => {
  const { token } = useAuth();
  const queryClient = useQueryClient();
  const {
    data: timeouts,
    isError,
    isFetching,
  } = useQuery(["timeouts", token], async () => getTimeouts(token ?? ""), { retry: 0 });

  const onSave = async (data: Timeouts) => {
    setTimeouts(data, token ?? "")
      .then(() => {
        queryClient.invalidateQueries("timeouts");
      })
      .catch(alert);
  };

  if (isFetching) {
    return <Header>Fetching</Header>;
  }

  if (isError) {
    return <Header>Error while fetching timeouts (maybe you need authorization)</Header>;
  }

  return (
    <div className="w-[100%] h-[100%] flex flex-col mt-[1.5rem]">
      <TimeoutsUI data={timeouts!} onSave={onSave} />
    </div>
  );
};

export default TimeoutsPage;
