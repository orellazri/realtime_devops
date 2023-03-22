import { stopAll } from "../../../utils/services";
import type { StartServicesRequest } from "../../../utils/types";

const id = 0;
const containerNames: string[] = [];

export async function POST({ request }) {
  const req: StartServicesRequest = await request.json();

  // const name = `sensor_${id}`;
  // startSensor(name, urls.kafka.local);
  // id++;
  // containerNames.push(name);

  return new Response("OK");
}

export async function DELETE() {
  stopAll(containerNames);
  return new Response("OK");
}
