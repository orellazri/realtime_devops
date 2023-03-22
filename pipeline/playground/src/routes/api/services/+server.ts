import { urls } from "../../../utils/constants";
import { startCompute, startReceiver, startSensor, stopAll } from "../../../utils/services";
import type { StartServicesRequest } from "../../../utils/types";

let id = 0;
const containerNames: string[] = [];

export async function POST({ request }) {
  const req: StartServicesRequest = await request.json();

  let name = `sensor_${id}`;
  containerNames.push(name);
  id++;
  startSensor(name, urls.kafka[req.sensor.send]);

  name = `compute_${id}`;
  containerNames.push(name);
  id++;
  startCompute(name, urls.kafka[req.compute.receive], urls.rabbitmq[req.compute.send]);

  name = `receiver_${id}`;
  containerNames.push(name);
  id++;
  startReceiver(name, urls.rabbitmq[req.receiver.receive]);

  return new Response("OK");
}

export async function DELETE() {
  stopAll(containerNames);
  return new Response("OK");
}
