export type Task = {
  id: number;
  expression: string;
  result: number;
  status: TaskStatus;
  createTime: number;
};

export enum TaskStatus {
  pending = "pending",
  executing = "executing",
  success = "success",
  error = "error",
}

export type Agent = {
  id: number;
  ping: number;
  active: boolean;
  maxThreads: number;
  runningThreads: number;
};

export type Timeouts = {
  add: number;
  sub: number;
  mul: number;
  div: number;
};

export type ErrorResp = {
  reason: string;
};

export type RegisterReq = {
  login: string;
  password: string;
};

export type LoginReq = RegisterReq;

export type LoginResp = {
  token: string;
};

export type GetTimeoutsResp = {
  timeouts: Timeouts;
};

export type SetTimeoutsReq = Partial<Timeouts>;

export type SetTimeoutsResp = {
  timeouts: Timeouts;
};

export type GetTaskResp = {
  task: Task;
};

export type GetTasksResp = {
  tasks: Task[];
};

export type SetTaskReq = {
  expression: string;
};

export type SetTaskResp = GetTaskResp;

export type GetAgentsResp = {
  agents: Agent[];
};
