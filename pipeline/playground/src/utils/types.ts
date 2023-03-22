export enum CommType {
  local = "local",
  cloud = "cloud"
}

export type CreateServicesRequest = {
  sensor: { send: CommType };
  compute: { send: CommType; receive: CommType };
  receiver: { receive: CommType };
};
