import express from "express";
import speakersRouter from "./speakersRouter";
const app = express();

// This middleware allows us to read body of requests.
// It's the equivalent of `app.use(bodyParser.json())`.
app.use(express.json());

// Initialize the routes.
app.use("/", speakersRouter);

export default app;
