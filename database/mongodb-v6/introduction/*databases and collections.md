# Databases and Collections

MongoDB stores data records as documents (BSON document) which are gathered together in collections. A database stores one or more collections of documents.

## Databases

In Mongodb, databases hold one or more collections of documents.

## Collections

MongoDB stores documents in collections. Collections are anologous to tables in relational databases.

**Explicit Creation**

MongoDB provides the db.createCollection() method to explicitly create a collection with various options, such as setting the maximum size or the documentation validation rules. If you are not specifying these options, you do not need to explicitly create the collection since MongoDB creates new collections when you first store data for the collections.

To modify these collection options, see collMod.

**Document Validation**

**Unique Identifiers**

Collections are assigned an immutable UUID. The collection UUID remains the same across all members of a replica set and shareds in a shared cluster.
