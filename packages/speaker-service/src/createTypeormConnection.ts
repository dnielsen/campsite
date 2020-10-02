import { createConnection, getConnectionOptions } from "typeorm";
import config from "./config";

export default async () => {
  const options = await getConnectionOptions(config.NODE_ENV);
  return createConnection({ ...options, name: "default" });
};
