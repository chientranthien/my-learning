# gRPC learning

## Quick start

- link: https://grpc.io/docs/languages/go/quickstart/
- I've learnt how to create a simple gRPC server and client, and how to generate the gRPC code from `.proto` file

## Basic turorial

### Why gRPC

- cross platform: can be generated to run in many evironment
- protobuff: efficiency serialization

### Basic gRPC

- I've learnt how to create a single RPC, client-side streaming, server-side streaming, bidirection

## Guides

### Authentication

- gRPC support
  - SSL/TLS
  - Google Cloud's ALTS
  - Token-based authentication with Google

### Benmarking

- gRPC has a Grafana dashboard to show their benmarking about
  - Latency
  - QPS
  - Scalability: Number of messages/second per server core
- Hardward: most instances has 8 core. For C++ and Java they additional suport QPS testing on 32 cores system
