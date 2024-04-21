import { BASE_URL } from ".";
import { ErrorResp, GetTimeoutsResp, SetTimeoutsReq, SetTimeoutsResp, Timeouts } from "./types";

export const getTimeouts = async (token: string): Promise<Timeouts> => {
  const url = BASE_URL + "/timeouts";
  const resp = await fetch(url, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }

  const res = (await resp.json()) as GetTimeoutsResp;
  return res.timeouts;
};

export const setTimeouts = async (timeouts: Partial<Timeouts>, token: string): Promise<Timeouts> => {
  const url = BASE_URL + "/timeouts";
  const body: SetTimeoutsReq = timeouts;
  const resp = await fetch(url, {
    method: "PATCH",
    headers: {
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  });

  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }

  const res = (await resp.json()) as SetTimeoutsResp;
  return res.timeouts;
};
