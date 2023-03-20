import { urls } from "../../../utils/constants";
import { startSensor, stopAll } from "../../../utils/services";

type Service = {
  type: "sensor" | "compute" | "receiver";
};

type CreateServicesRequest = {
  services: Service[];
};

let id = 0;
const containerNames: string[] = [];

export async function POST({ request }) {
  const req: CreateServicesRequest = await request.json();

  const name = `sensor_${id}`;
  startSensor(name, urls.kafka.local);
  id++;
  containerNames.push(name);

  return new Response("OK");
}

export async function DELETE() {
  stopAll(containerNames);
  return new Response("OK");
}
