import express from "express";
const app = express();
const port = 8000;

app.get("/", (req, res) => {
  res.send("Hello, KeyMesh backend!");
});

app.listen(port, () => {
  console.log("listening at http://localhost:" + port);
});

