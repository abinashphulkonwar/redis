import fetch from "node-fetch";
import { GetConfig } from "./init";
import { ErrorInterface, RedisClientError } from "./error";

export const Get = async ({
  Key,
  Command,
}: {
  Key: string;
  Command: string;
}) => {
  const config = GetConfig();
  const request = await fetch(
    `${config.url}/api/query/GET?command=${Command} ${Key}`
  );

  if (request.ok) {
    const res = (await request.json()) as unknown as {
      status: string;
      data: string;
      Is_LIST?: boolean;
    };
    return res;
  }
  const res = (await request.json()) as unknown as ErrorInterface;
  throw new RedisClientError(res);
};
