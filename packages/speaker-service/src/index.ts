import config from "./config";
import * as http from "http";
import app from "./app";
import createTypeormConnection from "./createTypeormConnection";

(async () => {
  // Connect to the database.
  await createTypeormConnection();

  const server = http.createServer(app);
  server.listen(config.PORT, () => {
    console.log(`Listening on port ${config.PORT}`);
  });
})();
