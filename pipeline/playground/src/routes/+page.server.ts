import type { CommType, StartServicesRequest } from "../utils/types";

export const actions = {
  start: async ({ request, fetch }) => {
    try {
      const data = await request.formData();

      const sensor_send = data.get("sensor_send") as CommType;
      const compute_send = data.get("compute_send") as CommType;
      const compute_receive = data.get("compute_receive") as CommType;
      const receiver_receive = data.get("receiver_receive") as CommType;

      const payload: StartServicesRequest = {
        sensor: { send: sensor_send },
        compute: { send: compute_send, receive: compute_receive },
        receiver: { receive: receiver_receive }
      };

      await fetch("/api/services", { method: "POST", body: JSON.stringify(payload) });
    } catch (e) {
      console.error("Failed to send request to API: " + e);
    }
  }
};
