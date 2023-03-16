import { exec } from "child_process";

export const runCommand = (cmd) => {
  exec(cmd, (error, stdout, stderr) => {
    if (error) {
      console.log(`ERROR: ${error.message}`);
      return;
    }
    if (stderr) {
      console.log(`ERROR (stderr): ${stderr}`);
      return;
    }
    console.log(`SUCCESS: ${stdout}`);
  });
};

export const startSensor = (name, url) => {
  runCommand(`docker run --rm -d --name ${name} \
    --network host \
    -e KAFKA_URL=${url} \
    pipeline-sensor`);
};
