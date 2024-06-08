# Pagination

- How pagination allows users to consume large data sets in bite-sized chunks
- The specific size of each page of data
- How API services should indicate that paging is complete
- How to define the page token format
- How to page across large chunks of data side a single resource

Rather than expecting all of the data in a single response interface, we'll enter a back-and-forth sequence whereby we request a small chunk of data at a time, iterating until there is no more data to consume.

How and where do we split the data

The term pagination comes from the idea that we want to page through the data, consuming data records in chunks, just like flipping through pages in a book. This is allows a consumer to ask for one chunk at a time and the API responds with the corresponding chunk, along with a pointer for how the consumer can retrieve the next chunk.

To accomplish this, our pagination pattern will rely on the idea of a cursor, using opaque page tokens as a way of expressing the loose equivalent of page numbers. Given this opaque, it will be important that the API responds both with the chunk of results as well as the next token. 

At a high level, we really want to way of picking up where we left off, effectively viewing a specific window of the available data. To do this, we'll rely on three different fields to convey our intent:

1. pageToken, which represents an opaque identifier, meaningful only to the API server of how to continue a previously started iteration of results.
2. maxPageSize, which allows the consumer to express a desire for no more than a certain number of results in a given response.
3. nextPageToken, which the API server uses to convey how the consumer should ask for more results with an additional request.


