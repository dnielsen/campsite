import { Router } from "express";
import Speaker from "./entity/Speaker";
import { v4 as generateUUID } from "uuid";

const speakersRouter = Router();

speakersRouter.get("/", async (req, res) => {
  const speaker = await Speaker.find();

  return res.status(200).send(speaker);
});

speakersRouter.post("/", async (req, res) => {
  const speaker = await Speaker.create({
    id: generateUUID(),
    bio: req.body.bio,
    name: req.body.name,
    headline: req.body.headline,
    photo: req.body.photo,
  }).save();

  return res.status(201).send(speaker);
});

speakersRouter.get("/:id", async (req, res) => {
  const speaker = await Speaker.findOne(req.params.id, {
    relations: ["sessions"],
  });
  // When speaker === undefined.
  if (!speaker) return res.sendStatus(404);

  return res.status(200).send(speaker);
});

speakersRouter.put("/:id", async (req, res) => {
  await Speaker.update(req.params.id, {
    headline: req.body.headline,
    photo: req.body.photo,
    name: req.body.name,
    bio: req.body.bio,
  });

  return res.sendStatus(204);
});

speakersRouter.delete("/:id", async (req, res) => {
  await Speaker.delete(req.params.id);

  return res.sendStatus(204);
});

export default speakersRouter;
