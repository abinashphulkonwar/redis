import fetch from "node-fetch";
import { GetConfig } from "./init";
import { ErrorInterface, RedisClientError } from "./error";
import { commands } from "./types";

export const Get = async ({
  Key,
  Command,
}: {
  Key: string;
  Command: commands;
}) => {
  const config = GetConfig();
  const request = await fetch(
    `${config.url}/api/query/GET?command=${Command} ${Key}`,
    {
      method: "GET",
      headers: {
        "content-type": "application/json",
      },
    }
  );

  if (request.ok) {
    const res = (await request.json()) as unknown as {
      status: string;
      data: string;
      Is_LIST?: boolean;
    };
    console.log(res);
    return res;
  }
  const res = (await request.json()) as unknown as ErrorInterface;
  throw new RedisClientError(res);
};
