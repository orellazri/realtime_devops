from datetime import datetime
import os
import signal
import socket
import sys
import time

from confluent_kafka import Producer

from sensor import Sensor

if __name__ == "__main__":
    def handle_sigint(*args):
        print(args)
        producer.flush()
        exit(0)    

    kafka_url = os.environ.get("KAFKA_URL")
    if not kafka_url:
        print("KAFKA_URL env var not found", file=sys.stderr)
        exit(1)

    producer = Producer({"bootstrap.servers": kafka_url,
                         "client.id": socket.gethostname()})

    signal.signal(signal.SIGINT, handle_sigint)

    sensor = Sensor()

    while True:
        sensor.update()
        now = datetime.utcnow().isoformat()
        x, y = sensor.get_coords()
        producer.produce("pipeline", key="coords", value=f"{now},{x},{y}")
        time.sleep(0.1)
