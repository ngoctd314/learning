# Subject-Based Messaging

Fundamental, NATS is about publishing and listening for messages. Both of these depend heavily on Subjects.

**What is a Subject**

At its simplest, a subject is just a string of characters that form a name which the publisher and subscriber can use to find each other. It helps scope messages into streams or topics.

**Wildcards**

NATS provides two wildcards that can take the place of one or more elements in a dot-separated subject. Subscribers can use these wildcards to listen to multiple subjects with a single subscription but Publishers will always use a fully specified subject, without the willcard.

**Mixing Wildcards**

The wildcard * can appear multiple times in the same subject. Both types can be used as well. For example, *.*.east.> will receive time.us.east.atlanta

**Subject Tokens**

It is recommended to keep the maximum number of tokens in your subjects to a reasonable value of 16 tokens max.