import signal
import socket
import time
from datetime import datetime

from confluent_kafka import Producer

from sensor import Sensor

if __name__ == "__main__":
    def handle_sigint(*args):
        print(args)
        producer.flush()
        exit(0)

    signal.signal(signal.SIGINT, handle_sigint)

    sensor = Sensor()

    producer = Producer({"bootstrap.servers": "127.0.0.1:29092",
                         "client.id": socket.gethostname()})
    while True:
        sensor.update()
        now = datetime.utcnow().isoformat()
        x, y = sensor.get_coords()
        producer.produce("pipeline", key="coords", value=f"{now},{x},{y}")
        time.sleep(0.1)
