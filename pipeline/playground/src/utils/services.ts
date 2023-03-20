import { execSync } from "child_process";

function runCommand(cmd: string) {
  execSync(cmd, (error: Error, stdout: Error, stderr: Error) => {
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
}

export function startSensor(name: string, kafka_url: string) {
  runCommand(`docker run --rm -d --name ${name} \
    --network host \
    -e KAFKA_URL=${kafka_url} \
    pipeline-sensor`);
}

export function startCompute(name: string, kafka_url: string, rabbitmq_url: string) {
  runCommand(`docker run --rm -d --name ${name} \
    --network host \
    -e KAFKA_URL=${kafka_url} \
    -e RABBITMQ_URL="${rabbitmq_url}" \
    pipeline-compute`);
}

export function stopAll(names: string[]) {
  runCommand(`docker stop ${names.join(" ")}`);
}
