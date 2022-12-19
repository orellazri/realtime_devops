import socket

from confluent_kafka import Producer

if __name__ == "__main__":
    producer = Producer({"bootstrap.servers": "127.0.0.1:29092",
                         "client.id": socket.gethostname()})
    for i in range(20):
        producer.produce("pipeline", key="mykey", value=str(i))
    producer.flush()
