import { BASE_URL } from ".";
import { Agent, ErrorResp, GetAgentsResp } from "./types";

export const getAgents = async ({ limit, offset }: { limit?: number; offset?: number }): Promise<Agent[]> => {
  const queryParams = new URLSearchParams();
  if (limit !== undefined) {
    queryParams.set("limit", `${limit}`);
  }
  if (offset !== undefined) {
    queryParams.set("offset", `${offset}`);
  }

  const url = BASE_URL + "/agents?" + queryParams.toString();
  const resp = await fetch(url);
  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }

  const res = (await resp.json()) as GetAgentsResp;
  return res.agents;
};
