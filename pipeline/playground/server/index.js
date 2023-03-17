import express from "express";
import { destroyAll, startSensor } from "./containers.js";

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

app.get("/destroy", (req, res) => {
  destroyAll(names);

  res.send("OK");
});

app.listen(3000);
