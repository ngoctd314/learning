# Documents

MongoDB stores data records as BSON documents. BSON is a binary representation of JSON documents, though it contains more data types that JSON.

## Document Limitations

Documents have the following attributes

**Document Size Limit**

The maximum BSON document size is 16 metabytes

The maximum document size helps ensure that a single document cannot use excessive amount of RAM or, during transmission, excessive amount of bandwidth. To store documents larger than the maximum size, MongoDB provides the GridFS API.

**Document Field Order**

Unlike JavaScript objects, the fields in a BSON document are ordered.

**Field Order in Queries**

For queries, the field order behavior is as follows:

- When comparing documents, field ordering is significant. For example, when comparing documents with fields a and b in a query

    - {a: 1, b: 1} is equal to {a: 1, b: 1}
    - {a: 1, b: 1} is not equal to {b: 1, a : 1}

- For efficient query execution, the query engine may reorder fields during query processing. Reordering fields may occur when processing these projection operations: $project, $addFields, $set and $unset.

    - Field reordering may occur in intermediate results as well as the final results returned by a query
    - Because some operations may reorder fields, you should not rely on specific field reordering in the results returned by a query that uses the projection operations listed earlier

**Field Order in Write Operations**

For write operations, MongoDB preserves the order of the document fields except for the following cases:

- The _id field is always the first field in the document
- Updates that including renaming of field names may result in the reordering of fields in the document 