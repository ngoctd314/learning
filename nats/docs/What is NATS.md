# What is NATS

Software applications and services need to exchange data.
NATS is an infrastructure that allows such data exchange, segmented in the form of messages. We call this a "message oriented middleware".

With NATS, application developers can:
- Effortlessly build distributed and scalable client-server applications.
- Store and distribute data in realtime in a general manner. This can flexibly be achieved cross various environments, languages, cloud providers and on-premises systems.

## NATS Client Applications

Developers use one of the NATS client libraries in their application code to allow them to publish, subscribe, request and reply between instances of the application or between completely separate applications.

## NATS Service Infrastructure

The NATS services are provided by one or more NATS server processes that are configured to interconnect with each other. The NATS service infrastructure can scale from a single NATS server process to a public global super-cluster of many clusters spanning all major cloud providers.

## Connecting NATS Client applications to the NATS servers

1. URL. A NATS URl
2. Authentication (if needed). NATS supports multiple authentication schemes (username/password, decentralized JWT, token, TLS certificates and Nkey with challenge).

## NATS Quality of service (QoS)

- At most once QoS: Core NATS offers an at most once quality of service. If a subscriber it not listening on the subject (no subject match), or is not active when the message is sent, the message is not received. This is the same level of guarantee that TCP/IP provides. Core NATS is a fire-and-forget messaging system. It will hold messages in memory and will never write messages directly to disk.
- At-least / exactly once QoS: NATS JetStream is built in to the NATS server (but needs to be enabled).