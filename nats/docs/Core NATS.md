# Core NATS

What is referred to as 'Core NATS' is the base set of functionalities and qualities of service offered by a NATS service infrastructure where none of the nats-server instances are configured to enable JetStream.

The 'Core NATS' functionalities are publish/subscribe with subject-based-addressing and queuing, with 'at-most-once' quality of service

## Publish-Subscribe

NATS implements a publish-subscribe message distribution model for one-to-many communication. A publisher sends a message on a subject and any active subscriber listening on that subject receives the message. Subscribers can also register interest in wildcard subjects that work a bit like a regular expression (but only a bit). This one-to-many pattern is sometimes called a fan-out.

Messsage are composed of

```go
type messasge struct {
    subject string
    payload []byte // max-size
    header map[string]any
    reply any
}
```

Messages have a maximum size (which is set in the server configuration with max_payload). The size is set to 1MB by default, but can be increased up to 64MB if needed.

### Pub/Sub Walkthrough

NATS is a publish subscribe messaging system based on subjects. Subscribers listening on a subject receive messages published on that subject. If the subscriber is not actively listening on the subject, the message is not received.

#### NATS Pub/Sub Walkthrough

This simple walkthrough demonstrates some ways in which subscribers listen on subjects and publishers send messages on specific subjects.

**1. Create Subscriber 1**

```bash
nats sub <subject>
```

**2. Create a Publisher and publish a message**

In another shell or command prompt, create a NATS publisher and send a message

```bash
nats pub <subject> <message>
```

## Request-Reply

Request-Reply is a common pattern in modern distributed systems. A request is sent, and the application either waits on the response with a certain timeout, or receives a response asynchronously.

The increased complexity of modern systems necessitates features like location transparency, scale-up and scale-down, observability.

**NATS makes Request-Reply simple and powerful**

- NATS supports the Request-Reply pattern using its core communication mechanism - publish and subscribe. A request is published on a given subject using a reply subject. Responders listen on that subject and send responses to the reply subject. Reply subjects are called "inbox". There are unique subjects that are dynamically directed back to the requester, regardless of the location of either party.

- Multiple NATS responders can form dynamic queue groups.

- NATS applications "drain before exiting" (processing buffered messages before closing the connection). This allows applications to scale down without dropping requests.

- The power of NATS even allows multiple responses

### Request-Replay Walkthrough

## Queue Groups 

When subscribers register themselves to receive messages from a publisher, the 1:N fan-out pattern of messaging ensures that any message sent by a publisher, reaches all subscribers that have registered. NATS provides an additional feature named "queue", which allows subscribers to register themeselves as part of a queue. Subscribers that are part of a queue, from the "queue group"

### How queue groups function

If a subscriber is registered based on a queue name, it will always receive messages it is subscribed to, based on the subject name. However, if more subscribers are added to the same queue name, they become a queue group, and only one randomly chosen subscriber of the queue group will consume a message each time a message is received by the queue group. Such distributed queues are a built-in load balancing feature that NATS provides.

**Advantages**

- Ensure application fault tolerance
- Workload processing can be scaled up or down
- No extra configuration required
- Queue groups can defined by the applition and their queue subscribers, rather than the server configuration

**Stream as a queue**

