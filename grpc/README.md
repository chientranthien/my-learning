# gRPC learning

## Following up

- How to generate tls cert files

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

### Cancellation

- gRPC support cancel function that client can send a cancel request to server. So server can propogate the cancel to its sub-call

### Compression

- gRPC can support encoding and client can specify it via the endocing header

### Custom metric

- gRPC record some basic metric. But you can add more of your own custom metrics

### Custom load balancer

- gRPC has it own client load balancer. But you can custom your load balancer

### Custom Name resolver

- Beside using the default DNS machanism. You can add your own custom Name resolver. The reason you want to do this because DNS only return IP. But if you need to watch for the new change you should has a custom name resolver to do it

### Load balancing

- Client side: load balancing, clients know the IPs address and make the call to servers directly. We can use this solution if we want to reduce latency for trusted clients
- Server side load balancing: we can use a L3/L4/L7 load balancer. So it can reduce the complicated on client side, but it will increase latency
- Look-aside load balancing: client subcribe to the load balancer to know about the server IPs and if there is any updates. But clients will send the requests directly to servers

### gRPC Client load balancing

- you can archived that by create a DNS to load balance or can create a custom service resolver implmenetion in code

### gRPC Server load balancing with nginx

- It's very easy to config nginx to load balance for grpc. You can config the upstream (grpc servers addresses) and the server to listen to a port. And that's all it need to config nginx with grpc

### Authentication

- Today, I've learnt that grpc can also support authentication via Authorization header. And Authentication must be used together with TLS/SSL

### Authorization

- gRPC can support authorization, you can specify authz policy for each method by method name (using)

### Compression

- gRPC can support compression like GZIP

### Cancellation

- you can cancel the request, and propogate it to the server

### gRPC

- you can create a health check, and gRPC will not call unhealthy instances
