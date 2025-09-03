import express from "express";
const app = express();
const port = 8000;

app.get("/hello", (req, res) => {
  res.send("hello world");
});

app.listen(port, () => {
  console.log("listening at http://localhost:" + port);
});

