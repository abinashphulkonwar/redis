import fetch from "node-fetch";
import { GetConfig } from "./init";
import { ErrorInterface, RedisClientError } from "./error";
import { commands } from "./types";

export const Write = async (event: {
  Data: {
    value: string;
    EX?: number;
  };
  Key: string;
  Commands: commands;
  IfNotExist: boolean;
}) => {
  const config = GetConfig();
  const request = await fetch(`${config.url}/api/write/add`, {
    method: "POST",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify(event),
  });

  if (request.ok) {
    const res = (await request.json()) as unknown as {
      body: {
        Value: string;
        EX: number;
        Type: string;
      };
      status: string;
    };
    console.log(res);

    return res;
  }
  const res = (await request.json()) as unknown as ErrorInterface;
  console.log(res);
  throw new RedisClientError(res);
};
