import tracing from "@opencensus/nodejs";
import { ZipkinTraceExporter } from "@opencensus/exporter-zipkin";
import jsonServer from "json-server";

const DB_PATH = "db.json";
const PORT = 4444;

const zipkinOptions = {
  url: "http://localhost:2020:9411/api/v2/spans",
  serviceName: "server",
};

const exporter = new ZipkinTraceExporter(zipkinOptions);
tracing.registerExporter(exporter).start();

const server = jsonServer.create();
const router = jsonServer.router(DB_PATH);
server.use(router);
server.listen(PORT, () => {
  console.log(`JSON server is running on port ${PORT}`);
});
