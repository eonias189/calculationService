"use client";
import axios from "axios";
import { contract } from "./.././contract";

const URL = "http://127.0.0.1:8081/";

export const fetchTasks = async (): Promise<contract.Task[]> => {
    const resp = await axios.get(URL + "getTasks");
    return contract.GetTasksResp.fromObject(resp.data).tasks;
};

export const addTask = async (expression: string) => {
    const body = new contract.AddTaskBody().toObject();
    body.expression = expression;
    axios.post(URL + "addTask", body);
};

export const fetchAgents = async (): Promise<contract.AgentData[]> => {
    const resp = await axios.get(URL + "getAgents");
    return contract.GetAgentsResp.fromObject(resp.data).agents;
};

export const fetchTimeouts = async (): Promise<contract.Timeouts> => {
    const resp = await axios.get(URL + "getTimeouts");
    return contract.GetTimeoutsResp.fromObject(resp.data).timeouts;
};

export const setTimeouts = async (timeouts: contract.Timeouts) => {
    const body = new contract.SetTimeoutsBody().toObject();
    body.timeouts = {
        add: timeouts.add,
        substract: timeouts.substract,
        multiply: timeouts.multiply,
        divide: timeouts.divide,
    };
    axios.post(URL + "setTimeouts", body);
};
