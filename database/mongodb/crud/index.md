# MongoDB CRUD operations

CRUD operations create, read, update, and delete documents

## Create Operations

Create or insert operations add new documents to a collection. If the collection does not currently exist, insert operations will create the collection

```go
db.collection.insertOne()
db.collection.insertMany()
```
Insert operations target a **single collection.** All write operations in MongoDB are **atomic** on the level of a single document

```go
db.users.insertOne({ // <- collection
    // document
    { 
        name: "sue",
        age: 26,
        state: "pending"
    }
})
```

## Read Operations

Read operations retrieve documents from a collection

```go
db.collection.find()
```
You can specify query filter

```go
db.collection.find(
    { age: {$gt: 18} },
    { name: 1, address: 1 }
).limit(5)
```

## Update operations

Update operations modify existing documents in a collection.

```go
db.collection.updateOne()
db.collection.updateMany()
db.collection.replaceOne()
```

You can specify criteria or filters, that identify the documents to update.

## Delete operations

Delete operations remove documents from a collection.

```go
db.collection.deleteOne()
db.collection.deleteMany()
```
Delete operations target a single collection. All write operations in MongoDB are atomic on the level of a single document.