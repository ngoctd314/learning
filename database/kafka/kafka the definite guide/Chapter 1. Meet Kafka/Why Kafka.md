# Why Kafka?

## Multiple Producers, Consumers

## Disk-Based Retention

Not only Kafka handle multiple consumers, but durable message retention means that consumers do not always need to work in real time. Messages are written to disk and will be stored with configurable retention rules. These options can be selected on a per-topic basis, allowing for different streams of messages to have different amounts of retention depending on the consumer needs. Durable retention means that if a consumer falls behind, either due to slow processing or a burst in traffic, there is no danger of losing data. It also means that maintenance can be performed on consumers, taking applications offline for a short period of time, with no concern about messages backing up the producer or getting lost. Consumers can be stopped, and the messages will be retained in kafka. This allows them to restart and pick up processing message where they left off with no data losts. 

## Scalable

## High Performance


