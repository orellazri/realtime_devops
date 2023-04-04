# DevOps Infrastructure for Realtime Applications

[![Test](https://github.com/orellazri/realtime_devops/actions/workflows/test.yml/badge.svg)](https://github.com/orellazri/realtime_devops/actions/workflows/test.yml)

Nowadays, most realtime applications use a monolithic approach in regard to deployment, monitoring, and transporting information. While this is a proven approach, we are trying to test if using a modern one, with the latest DevOps tools and principles can benefit said applications.

We will benchmark the performance of different approaches to sending data, such as
basic transport, streaming to brokers, and processing the information mid-transport.

This repository will contain the code for the performance benchmarks, and everything else this project might include.

## Directories

- **performance**: Benchmarking different tools and protocols and generating charts
- **pipeline**: Testing the performance of a simple pipeline consisting of several microservices, running locally/in the cloud
- **playground**: Modular playground for testing different scenarios of communicators talking via configurable protocols, addresses, etc.
- **stacks**: Docker compose files needed for starting all relevant services locally
