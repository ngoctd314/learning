# How the internet works

## The internet protocol suite

The network engineers developed the TCP to ensure a reliable exchange of information between computers. When a computer sends a message to another machine via TCP, the message is split into data packets that are sent toward their eventual destination address. The computers that make up the internet push each packet toward the destination without having to process the whole message.

Once the recipient computer receives the packets, it assembles them back into a usable order according to the sequence number on each packet. Every time the recipient receives a package, it sends a receipt. If the recipient fails to ack receipt of a packet, the sender resends that packet, posibly along a different network path.

Packets are now sent with a checksum that allows recipients to detect data corruption and determine whether packets need to be resent.

TCP remains the common protocol because of its delivery guarantees, but nowadays, serveral other protocols are also used over the internet. The UDP, for instance, is a newer protocol that deliberately allows packets to be dropped so that data can be streamed at a constant rate. UDP is commonly used for streaming live video, since consumers prefer a few dropped frames over having their feed delayed when the network gets congested.

### Internet Protocol Addresses

Data packets on the internet are sent to IP addresses, num-bers assigned to individual internet-connected computers. Regional authorities then grant the blocks of addresses to internet service providers (ISPs) and hosting companies within their region. When you connect your browser to the internet, your ISP assigns your computer an IP addresses that stays fixed for a few months.

### The Domain Name System

Domain Names are simply placeholders for IP addresses. Domain names, like IP addresses, are unique, and have to be registered before use with private organizations called domain registrars.

When browsers encounter a domain name for the first time, they use a local domain name server (typically hosted by an ISP) to look it up, and then cache the result to prevent time-consuming lookups in the future. This caching behavior means that new domains or changes to existing domains take a while to propagate on the internet.

DNS caching enables a type of attack called DNS poisoning, whereby a local DNS cache is deliberately corrupted so that data is routed to a server controlled by an attacker.

### Application Layer Protocols

TCP allows two computers to reliably exchange data on the internet, but it doesn't dictate how the data being sent should be interpreted. For that to happen, both computers need to argree to exchange information through another. Protocols that build on top of TCP of UDP are called application layer protocols.

For example, emails sent using the SMTP, file servers make downloads available via the File Transfer Protocol (FTP), and web servers use HTTP.

| Layer             | Protocol                                   |
| ----------------- | ------------------------------------------ |
| Application Layer | DNS, FTP, HTTP, IMAP, POP, SMTP, SSH, XMPP |
| Transport layer   | TCP, UDP                                   |
| Internet layer    | IPv4, IPv6                                 |
| Network layer     | ARP, MAC, NDP, OSPF, PPP                   |

### HyperText Transfer Protocol

Web servers use the HyperText Transfer Protocol (HTTP) to transport web pages and their resources to user agents such as web browsers. In an HTTP conversation, the user agent generates requests for particular resources. Web servers, expecting these requests, return responses containing either the requested resource, or an error code if the request can't be fulfilled. Both HTTP requests and responses are plaintext messages, though they're often sent in compressed and encrypted form.

**1. HTTP Requests**

**Method** Also known as a verb, this describes the action that the user agent wants the server to perform.

**Universal resource locator (URL)** This describes the resource being manipulated or fetched

**Headers** These supply metadata such as the type of content the user agent is expecting or whether it accepts compressed responses.

**Body** This optional component contains any extra data that needs to be sent to the server.

```txt
GET http://example.com/
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS x 10_13_6)
AppleWebKit/537.36 (KHTMl, like Gecko) Chrome...
Accept: text/html,application/xhtml+xml,application/xml; */*
```

**2. HTTP Responses**

HTTP responses sent back by a web server begin with a protocol description, a three-digit status code, and typically, a status message that indicates whether the request can be fulfilled.

```txt
Content-Encoding: gzip
Accept-Ranges: bytes
Cache-Control: max-age=604800
Content-Type: text/html
Content-Length: 606
```

**3. Stateful Connections**

Web servers typically deal with many user agents at once, but HTTP does nothing to distinguish which requests are coming from which user agent. In nowadays, HTTP conversations need to be made stateful. A connection or conversation between a client and a server is stateful when they perform a "handshake" and continue to send packets back and forth until one of the communicating parties decides to terminate the connection.

When a web server wants to keep track of which user it's responding to with each request, and thus achieve a stateful HTTP conversation, it needs to establish mechanism to track the user agent as it makes the subsequent requests. The entire conversation between a particular user agent and a web server is called an HTTP session. The most common way of tracking sessions is for the server to send back a Set-Cookie header in the initial HTTP response. This asks the user agent receiving the response to store a cookie, a small snippet of text data pertaining to that particular web domain.

Session information contained in cookies is a juicy target for hackers. If an attacker steals another user's cookie, they can pretend to be that user on the website. Similarly, if an attacker successfully persuades a website to accept a forged cookie, they can impersonate any user they please.

**4. Encryption**

When the web was first invented, HTTP requests and responses were sent in plaintext form, which meant they could be read by anyone intercepting the data packets; this kind of interception is known as man-in-the-middle attack.

To secure their communications, web servers and browsers send requests and responses by using Transport Layer Security (TLS), a method of encryption that provides both privacy and data integrity. TLS ensures that packets intercepted by a third party can't be decrypted without the appropriate encryption keys.

HTTP conversations conducted using TLS are called HTTP Secure (HTTPS). HTTPS requires the client and server to perform a TLS hand-shake in which both parties agree on an encryption method (a cipher) and exchange encryption keys. One the hand-shake is complete, any futher messages (both requests and responses) will be opaque to outsides.

### Summary

TCP enables reliable communication between internet-connected computers that each have an IP address. The Domain Nam System provides human readable aliases for IP addresses. HTTP builds on top of TCP to send HTTP requests from user agents (such as web browsers) to web servers, which in turn reply with HTTP responses. Each request is sent to a specific URL, and you learned about various types of HTTP methods. Web servers respond with status codes, and send back cookies to initiate stateful connections. Finally, encryption (in the form of HTTPS) can be used to secure communication between a user agent and a web server.
