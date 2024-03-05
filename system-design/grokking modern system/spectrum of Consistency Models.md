# Spectrum of Consistency Models

## What is consistency?

In distributed systems, consistency may mean many things: One is that each replica node has the same view of data at a given point in time. The other is that each read request gets the value of the recent write. These are not the only definitions of consistency, since there are many forms of consistency. Normally, consistency models provide us with abstractions to reason about the correctness of a distributed system doing concurrent data reads, writes, and mutations.

The two ends of the consistency spectrum are:

- Strongest consistency
- Weakest consistency

Eventual consistency -> Causal consistency -> Sequential consistency -> Strict consistency/linearizability

(Weakest consistency model)                                             (Strongest consistency model)

There is a different between consistency in ACID properties and consistency in the CAP theorem. Database rules are at the heart of ACID consistency. If a schema specifies that a value must be unique, a consistent system will ensure that the value is unique throughout all actions. If a foreign key indicates that deleting one row will also delete associated rows, a consistent system ensures that the state can't contain related rows once the base row has been destroyed.

CAP consistency guarantees that, in a distributed system, every replica of the same logical value has the same precise value at all times. It's worth noting that this is a logical rather than a physical guarantee. Due to the speed of light, replicating numbers throughout a cluster may take some time. By preventing clients from accessing different values at separate nodes, the cluster can nevertheless give a logical picture.

## Eventual consistency

**Eventual consistency** is the weakest consistency model. The applications that don't have strict ordering requirements and don't require reads to return the latest write choose this model. Eventual consistency ensures that all the replicas coverage on a final value after a finite time and when no more writes are coming in. If new writes keep coming, replicas of an eventually consistent system might never reach the same state. Until the replicas converge, different replicas can return different values.

Eventual consistency ensures high availability:

The domain name system is a highly availability system that enables name lookups to a hundred million devices across the Internet. It uses an eventual consistency model and doesn't necessarily reflect the latest values.

## Causal consistency


## Sequential consistency

## Strict consistency aka linearizability

A strict consistency or linearizability is the strongest consistency model. This model ensures that a read request from any replicas will get the latest write value. Once the client receives the ack that the write operation has been performed, other clients can read that value.

Linearizability is challenging to achieve in a distributed system. Some of the reasons for such challenges are variable network delays and failures. The following slides show depicts how variable network delays make it possible for different parties to see different values.

Usually, synchronous replication is one of the ingredients for achieving strong consistency, though it in itself is not sufficient. We might need consensus algorithm such as Paxos and Raft to achieve strong consistency.

Linearizability affects the system's availability, which is why it's not always used. Applications with strong consistency requirements use techniques like quorum-based replication to increase the system's availability.

**Example**

Updating an account's password requries strict consistency. For example, if we suspect suspicious activity on our bank account, we immediately change our password so that no unauthorized users can access our account. If it were possible to access our account an old password due to a lack of strict consistency, then changing passwords would be a useless security strategy.

## Summary

Linearizability services appear to carry out transactions/operations in sequential, real-time order. They make it easier to create suitable applications on top of them by limiting the number of values that services can return to application processes.

Application programmers have to compromise performance and availability if they use services with strong consistency models. The models may break the invariants of applications built on top of them in exchange for increased performance.
