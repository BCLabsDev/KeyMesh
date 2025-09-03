import "./utils/info.js"
import 'dotenv/config'
import { logger } from "./utils/log.js";
import { InitDBClient } from "./orm/client.js";
import express from "express";

const app = express();
const port = 8000;

try {
    await InitDBClient()
} catch(error) {
    logger.error("初始化操作失败，报错:"+error)
    process.exit(1);
}

app.get("/", (req, res) => {
  res.send("Hello, KeyMesh backend!");
});

app.listen(port, () => {
  logger.info("Web Service listening at " + port);
});

