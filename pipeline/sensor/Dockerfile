FROM ubuntu:latest

ARG DEBIAN_FRONTEND=noninteractive

# Install python3 and pip
RUN apt-get update && apt-get install -y software-properties-common gcc && \
    add-apt-repository -y ppa:deadsnakes/ppa
RUN apt-get update && apt-get install -y python3.11 python3-pip

# Install requirements
ADD requirements.txt .
RUN pip install -r requirements.txt

ADD src src/
CMD ["python3", "src/main.py"]
