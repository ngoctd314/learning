# Update Documents

## Update Documents in a Collection

To update a document use $set to modify field values.

Some update operators, such as $set, will create the field if the field does not exist. 

## Behavior 

**Atomicity**

All write operations in MongoDB are atomic on the level of a single document.

**_id Field**

Once set, you cannot update the value of the _id field nor can you replace an existing document with a replacement document that has a different _id field value.

**Field Order**

For write operations, MongoDB preserves the order of the document fields except for:

- The _id field is always the first field in the document.
- Updates that include renaming of field names may result in the reordering of fields in the document.

**Write Acknowledgement**

With write concerns, you can specify the level of ack requested from MongoDB for write operations.