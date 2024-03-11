# Domain Name System

Learn how domain names get translated to IP addresses through DNS.

## The origins of DNS

Computers are uniquely identified by IP addresses - for example 104.18.2.119 is an IP address. We use IP addresses to visit a website  hosted on a machine.

We need a phone book-like repository that can maintain all mappings of domain names to IP addresses.

## What is DNS?

The domain name system (DNS) is the Internet's naming service that maps human-friendly domain names to machine-readable IP addresses. The service of DSN is transparent to users. When a user enters a domain name in the browser, the browser has to translate the domain name to IP address by asking the DNS infrastructure. Once the desired IP address is obtained, the user's request is forwarded to the destination web server.

## Important details

- Name servers: It's important to understand that the DNS isn't a single server. It's a complete infrastructure with numerous servers. DNS servers that respond to user's queries are called name servers.
- Resource records: The DNS database stores domain name to IP mappings in the form of resource records (RR). The RR is the smallest unit of information that users request from the name servers. There are different types of RRs. The table below describes common RRs. The three important pieces of information are type, name, and value. The name and value change depending upon the type of the RR.

- Caching: DNS uses caching at different layers to reduce request latency for the user. Caching plays an important role in reducing the burden on DNS infrastructure because it has to cater to the queries of the entire Internet.
- Hierarchy: DNS name servers are in a hierarchical form. The hierarchical structure allows DNS to be highly scalable because of its increasing size and query load. In the next lesson, we'll look at how a tree-like structure is used to manage the entire DNS database.

## How the Domain Name System Works

- How is the DNS hierarchy formed using various types of DNS name servers?
- How is caching performed at different levels of the Interent to reduce the querying burden over the DNS infrastructure?
- How does the distributed nature of the DNS infrastructure help its robustness?

## DNS hierarchy

As started before, the DNS isn't a single server that accepts requests and responds to user queries. It's a complete infrastructure with name servers at different hierarchies.

There are mainly four types of servers in the DNS hierarchy:

1. DNS resolver: Resolvers initiate the querying sequence and forward requests the the other DNS name servers. Typically, DNS resolvers lie within the premise of the user's network. However, DNS resolvers can also cater to user's DNS queries through caching techniques, as we will see shortly. These servers can also be called local or default server.
2. Root-level name servers: These servers receive requests from local servers. Root name servers maintain name servers based on top-level domain names such as .com, .edu, .us, and so on. For instance, when a user requests the IP address of educative.io, root-level name servers will return a list of top-level domain (TLD) servers that hold the IP addresses of the .io domain.
3. Top-level domain
4. Authoritative name servers

**Question:** How are DNS names processed? For example, will educative.io be processed from left to right or right to left?

Unlike UNIX files, which are processed from left to right, DNS names are processed from right to left. In the case of educative.io, the resolvers will first resolve the .io part, then educative, and so on.

## Iterative versus recursive query resolution

There are two ways to perform a DNS query:

1. Iterative: The local server requests the root, TLD, and the authoritative servers for the IP address.
2. Recursive: The end user requests the local server. The local server further requests the root DNS name servers...
