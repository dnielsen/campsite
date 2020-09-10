import { NextApiRequest, NextApiResponse } from "next";

const handler = (req: NextApiRequest, res: NextApiResponse) => {
  res.status(200).json([{ id: "1234", name: "john" }]);
};

export default handler;
