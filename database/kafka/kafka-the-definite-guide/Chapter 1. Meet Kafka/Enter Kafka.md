# Enter Kafka

Apache Kafka was developed as publish/subcribe messaging system designed to solve this problem. It is often described as a "distributed commit log" or more recently as a "distributing streaming platform". A filesystem or database commit log is designed to provide a durable record of all transactions so that they can be replayed to consistently build the state of a system. Similarly, data within Kafka is stored durably, in order, and can be read deterministically. In addition, the data can be distributed within the system to provide additional protections against failures, as well as significant opportunities for scaling performance.

## Messages and Batches

The unit of data within Kafka is called a message. You can think message as similar to row or a record. A message is simply an array of bytes as far as Kafka is concerned. A message can have an optional piece of metadata, has no specific meaning to Kafka. Keys are used when messages are to be written to partitions in a more controlled manner. Message with the same key are always written to the same partition.

For efficiency, messages are written into Kafka in batches. A batch is just a collection of messages, all of which are being produced to the same topic and partition. An individual round trip across the network for each message would result in excessive overhead, and collecting messages together into a batch reduces this. Of course, this is a trade-off between latency and throughput: the larger batches, the more messages that can be handled per unit of time, but the longer it takes an individual message to propagate. Batches are also typically compressed, providing more efficient data transfer and storage at the cost of some processing power. 

## Schemas

While messages are opaque byte arrays to Kafka itself , it is recommended that additional structure, or schema, be imposed on the message content so that it can be easily understood. There are many options available for message schema, depending on your application's individual needs. Simplistic systems, such as JavaScript Object Notation (JSON) and Extensible Markup Language (XML), are easy to use and human readable. However, they lack features such as robust type handling and compatibility between schema versions. Many Kafka developers favor the use of Apache Avro, which is serialization fw originally developed for Hadoop. Arvo provides a compact serialization format, schemas that are separate from the message payloads and that do not require code to be generated when they change, and strong data typing and shema evolution.

A consistent data format is important in Kafka, as it follows writing and reading messages to be decoupled. When these tasks are tightly coupled, applications that subscribe to messages must be updated to handle the new data format, in parallel with the old format.

## Topics and Partitions

Messages in Kafka are categorized into topics. The closest analogies for a topic are a database table or a folder in a filesystem. Topics are additionally broken down into a number of partitions. Going back to the "commit log" description, a partition is a single log. Messages are written to it in an append-only fashion and are read in order from beginning to end. Note that as a topic typically has multiple partitions, there is no guarantee of message ordering across the entire topic, just within a single partition. Partitions are also the way that Kafka provides redundancy and scalability. Each partition can be hosted on a different server, which means that a single topic can be scaled horizontally across multiple servers to provide performance far beyond the ability of a single server. Additionally, partitions can be replicated, such that different servers will store a copy of the same partition in case one server fails.

## Producers and Consumers

Kafak clients are users of the system, and there are two basic types: producers and consumers. There are also advanced client APIs - Kafka Connect API for data integration and Kafka Streams for stream processing. The advanced clients use producers and consumers as building blocks and provide higher-level functionality on top.

Producers create new messages. In other publish/subscribe systems, these may be called publishers or writers. A message will be produced to a specific topic. By default, the producer will balance message over all partitions of a topic evenly. This is typically done using the message key and a paritioner that will generate a hash of the key and map it to a specific partition. This ensures that all messages produced with a given key will get written to the same partition. The producer could also use a custom partitioner that follows other business rules for mapping messages to partitions.

Consumers read messages. In other publish/subscribe systems, these clients may be called subsribers or readers. The consumer subscribes to one or more topics and reads the messages in the order in which they were produced to each partition. The consumer keeps track of which messages it has already consumed by keeping track of the offset of messages. The offset - an integer value that continually increases - is another piece of metadata that Kafka adds to each message as it is produced. Each message in a given partition has a unique offset, and the following message has a greater offset (though not necessarily monotonically greater). By storing the next posible offset for each partition, typically in Kafka itself, a consumer can stop and restart without losing its place.

Consumers work as part of a consumer group, which is one or more consumers that work together to consume a topic. The group ensures that each partition is only consumed by one member. In this way, consumers can horizontally scale to consume topics with a large number of messages. Additionally, if a single consume fails, the remaining members of group will reassign the partitions being consumed to take over for the missing member.

In this way, consumers can horizontally scale to consume topics with a large number of message. Additionally, if a single consumer fails, the remaining members of the group will reassign the partitions being consumed to take over for the missing member. 

## Broker and Clusters

A single Kafka server is called a broker. The broker receives messages from producers, assigns offsets to them, and writes the messages to storage on disk. It also services consumers, responding of fetch requests for partitions and responding with the messages that have been published.

Kafka brokers are designed to operate as part of a cluster. Within a cluster of brokers, one broker will also function as the cluster controller (elected automatically from the live members of the cluster). The controller is responsible for administrative operations, including assigning partitions to brokers and monitoring for broker failures. A partition is owned by a single broker in the cluster, and than broker is called the leader of the partition. A replicated partition is assigned to addition brokers, called follower of the partition. All producers must connect to the leader in order to publish messages, but consumers may fetch from either the leader or one of the followers.

A key feature of Apache Kafka is that of retention, which is the durable storage of messages for some period of time. Kafka brokers are configured with a default retention setting by topics, either retaining messages for some period of time (e.g. 7 days) or until the partition reaches a certain size in bytes (e.g. 1GB). Once these limits are reached, messages are expired and deleted. Individual topics can also be configured with their own retention settings so that messages are stored for only as long as they are useful. For example, a tracking topic might be retained for several days, whereas application metrics might be retained for only a few hours.

## Multiple Clusters

As Kafka deployments grow, it is often advantageous to have multiple clusters. There are several reasons why this can be useful:

- Segregation of types of data
- Isolation for security requirements
- Multiple datacenters (disaster recovery)

When working with multiple datacenters in particular, it is often required that messages be copied between them. In this way, online applications can have access to user activity at both sites. For example, if a user changes public information in their profile, that change will need to be visible regardless of the datacenter in which search results are displayed. The replication mechanism within the Kafka clusters are designed only to work within a single cluster, not between multiple clusters.
