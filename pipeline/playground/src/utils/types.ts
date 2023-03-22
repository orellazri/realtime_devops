export enum CommType {
  local = "local",
  cloud = "cloud"
}

export type StartServicesRequest = {
  sensor: { send: CommType };
  compute: { send: CommType; receive: CommType };
  receiver: { receive: CommType };
};
