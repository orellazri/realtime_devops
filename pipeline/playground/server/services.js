import { execSync } from "child_process";

export function runCommand(cmd) {
  execSync(cmd, (error, stdout, stderr) => {
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

export function startSensor(name, url) {
  runCommand(`docker run --rm -d --name ${name} \
    --network host \
    -e KAFKA_URL=${url} \
    pipeline-sensor`);
}

export function stopAll(names) {
  runCommand(`docker stop ${names.join(" ")}`);
}
