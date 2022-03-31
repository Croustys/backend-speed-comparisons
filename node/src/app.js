import express from "express";

const PORT = process.env.PORT || 4567;

const app = express();

app.get("/", (req, res) => res.send("Server Working as intended"));
app.get("/ping", (req, res) => res.json({ message: "PONG" }));
app.get("/test", (req, res) => {
  const startTime = new Date();
  const total = 10000000000;
  let testSum = 0;
  for (let index = 0; index < total; index++) {
    testSum++;
  }
  if (testSum === total) {
    const endTime = new Date();
    const runTime = endTime - startTime;
    res.json({ complete: true, runTime: `${runTime} ms` });
  } else res.status(500).json({ error: "error" });
});

app.listen(PORT, () => {
  console.log(`Listening on http://localhost:${PORT}`);
});
