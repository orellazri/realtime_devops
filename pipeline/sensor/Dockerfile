FROM python:3.10

RUN apt-get update && \
    apt-get install -y --no-install-recommends gcc git libssl-dev g++ make && \
    cd /tmp && git clone https://github.com/edenhill/librdkafka && \
    cd librdkafka && git checkout tags/v2.0.2 && \
    ./configure && make && make install && \
    ldconfig &&\
    cd ../ && rm -rf librdkafka

# Install requirements
ADD requirements.txt .
RUN pip install -r requirements.txt

ADD src src/
CMD ["python3", "src/main.py"]
