import express from "express";
import { startSensor, stopAll } from "./services.js";

const app = express();
app.use(express.json());

// Global container names for services started
let names = [];

/*
  Start a service
  {
    data: [
      {
        type: ...
      }
    ]
  }
*/
app.post("/start", (req, res) => {
  const name = `sensor_${names.length}`;
  names.push(name);
  startSensor(name, "127.0.0.1:29092");

  res.send("OK");
});

app.get("/stop", (req, res) => {
  if (names.length > 0) {
    stopAll(names);
    names = [];
  }

  res.send("OK");
});

app.listen(3000);
