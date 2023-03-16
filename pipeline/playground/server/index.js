import express from "express";
import { startSensor } from "./containers.js";

const app = express();
app.use(express.json());

// Global id for services started
let id = 0;

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
  startSensor(`sensor_${id}`, "127.0.0.1:29092");
  id++;
  // for (let service of req.body.data) {
  // }

  res.send("OK");
});

app.listen(3000);
