import { BASE_URL } from ".";
import { ErrorResp, GetTasksResp, SetTaskReq, SetTaskResp, Task } from "./types";

export const getTasks = async (
  token: string,
  { limit, offset }: { limit?: number; offset?: number }
): Promise<Task[]> => {
  const queryParams = new URLSearchParams();
  if (limit !== undefined) {
    queryParams.set("limit", `${limit}`);
  }
  if (offset !== undefined) {
    queryParams.set("offset", `${offset}`);
  }

  const url = BASE_URL + "/tasks?" + queryParams.toString();
  const resp = await fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }

  const res = (await resp.json()) as GetTasksResp;
  return res.tasks;
};

export const addTask = async (expression: string, token: string): Promise<Task> => {
  const url = BASE_URL + "/tasks";
  const body: SetTaskReq = { expression };
  const resp = await fetch(url, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  });
  if (!resp.ok) {
    const res = (await resp.json()) as ErrorResp;
    throw new Error(res.reason);
  }

  const res = (await resp.json()) as SetTaskResp;
  return res.task;
};
