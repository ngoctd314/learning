# Chapter 2. Data Models and Query Languages

The limits of my language mean the limit of my world.

Data models are perhaps the most important part of developing software, because they have such a profound effect: not only on how the software is written, but also on how we think about the problem that we are solving.

Most applications are built by layering one data model on top of another. For each layer, the key question is: how is it represented in terms of the next-lower layer?

1. As an application developer, you look at the real world and model it in terms of objects or data structures, and APIs that manipulate those data structures.
2. When you want to store those data structures, you express them in terms of a general-purpose data model, such as JSON or XML documents, tables in a relational database, or a graph model.
3. The engineers who built your database software decided on a way of representing that JSON/XML/relation/graph data in terms of bytes in memory, on disk, or on a network. The representation may allow the data to be queries, searched, manipulated, and processed in various ways.
4. On yet lower levels, hardware engineers have figured out how to represent bytes in terms of electrical currents, pulses of light, magnetic fields, and more.

In a complex application there may be more intermediary levels, such as APIs build upon APIs, but the basic idea is still the same: each layer hides the complexity of the layers below it by providing a clean data model.

## Relational Model Versus Document Model

The best-known data model today is probably that of SQL, based on the relational model proposed by Edgar Codd in 1970.

## The Birth of NoSQL

There are several driving forces behind the adoption of NoSQL databases, including:

- A need for greater scalability that relation databases can easily achieve, including very large datasets or very high write throughput.
- A widespread preference for free and open source software over commercial database products.
- Specialized query operations that are not well supported by the relational model.
- Frustration with the restrictiveness of relational schemas, and a desire for a more dynamic and expressive data model.

## The Object-Relational Mismatch


