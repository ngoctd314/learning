# Interact with data in Redis

## Search and query

Searching and querying Redis data

Redis stack offers enhanced Redis experience via the following search and query features:

- A rich query language
- Incremental indexing on JSON and hash documents
- Vector search
- Full-text search
- Geospatial queries
- Aggregations

### Basic constructs

Basic constructs for searching and querying Redis data.

You can use Redis Stack as a powerful search and query engine. It allows you to create indexes and perform efficient queries on structured data, as well as text-based and vector searches on unstructured data.

**Documents**

A document is the basic unit of information. It can be any hash or JSON data object you want to be able to index and search. Each document is uniquely identifiable by its key name.

**Fields**

A document consists of multiple fields, where each field represents a specific attribute or property of the document.

**Indexing fields**

Not all fields are relevant to perform search operations, and indexing all fields may lead to unnecessary overhead. That's why you have the flexibility to choose which fields should be indexed for efficient search operations. By indexing a field, you enable Redis Stack to create an index structure that optimizes search performance on that field.

**Schema**

The index structure is defined by a schema. The schema defines how fields are stored and indexed. It specifies the type of each field and other important information.

### Schema definition

An index structure is defined by a schema. The schema specifies the fields, their types, whether they should be indexed or stored, and other additional configuration options.

## Transactions

Redis Transactions allow the execution of a group of commands in a single step, they are centered around the commands MULTI, EXEC, DISCARD and MATCH. Redis Transactions make two important guarantees:

- All the commands in a transaction are serialized and executed sequentially. A request sent by another client will never be served in the middle of the execution of Redis Transaction. This guarantees that the commands are executed as a single isolated operation.

## Redis Pub/Sub

SUBSCRIBE, UNSUBSCRIBE and publish implement the Publish/Subscribe messaging paradigmm where senders (publishers) are not programmed to send their messages to specific receivers (subscribers). Rather published messages are characterized into channels, without knowledge of what (if any) subscribers there may be. Subscribers express interest in one or more channels and only receive messages that are of interest, without knowledge of what (if any) publishers there are.

```txt
SUBSCRIBE channel11 ch:00
```

Messages sent by other clients to these channels will be pushed by Redis to all the subscribed clients. Subscribers receive the messages in the order that the messages are published.


