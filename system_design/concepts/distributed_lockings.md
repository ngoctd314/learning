# Distributed Locking

## Official Doc

Distributed locks are a very useful primitive in many environments where different processes must operate with shared resources in a mutually exclusive way.

There are a number of libraries and blog posts describing how to implement a DLM (Distributed Lock Manager) with Redis, but every library uses a different approach, and many use a simple approach with lower guarantees compared to what can be achieved with slightly more complex designs.

**Safety and Liveness Guarantees**

We are going to model our design with just three properties that, from our point of view, are the minimum guarantees needed to use distributed locks in an effective way.

1. Safety property: Mutual exclusion. At any given moment, only one client can hold a lock.
2. Liveness property A: Deadlock free. Eventually it is always possible to accquire a lock, even if the client that locked a resource crashes or gets partitioned.
3. Liveness property B: Fault tolerance. A long as the majority of Redis nodes are up, clients are able to acquire and release locks.

**Why Failover-based Implementations Are Not Enough**

The simplest way to use redis to lock a resource is to create a key in an instance. The key is usually created with a limited time to live, so that eventually it will get released. When the client needs to release the resource, it deletes the key.

This works well, but there is a problem: this is a single point of failure in our architecture. What happens if the Redis master goes down? Well, let's add a replica! Add use it if the master is unavailable. This is unfornately not viable. By doing so we can't implement our safety property of mutual exlusion, because Redis replication is asynchronous. 

There is a race condition with this model:

1. Client A accquires the lock in the master
2. The master crashes before the write to the key is transmitted to the replica
3. The replica gets promoted to master
4. Client B acquires the lock to the same resource A already holds a lock for. SAFETY VIOLATION!

Sometimes it is perfectly fine that, under special circumstances, for example during a failure, multiple clients can hold the lock at the same time. If this is the case, you can use your replication based solution. Otherwise we suggest to implement the solution described in this document.

**Correct Implementation with a Single Instance**

Before trying to overcome the limitation of the single instance setup describe above, let's check how to do it correctly in this simple case, since this is actually a viable solution in applications where a race condition from time to time is acceptable, and because locking into a single instance is the foundation we'll use for the distributed algorithm describe here.

```bash
SET resource_name my_random_value NX PX 30000
```

The command will set the key only if it does not already exist (NX option), with an expire of 30000 milliseconds (PX option). The key is set to a value "my_random_value". This value must be unique across all clients and all lock requests.

Basically the random value is used in order to release lock in a safe way, with a script that tells Redis: remove the key only if it exists and the value stored at the key is exactly the one I expect to be. This is accomplished by the following script:

```bash
if redis.call("get", KEYS[1]) == ARGV[1] then
    return redis.call("del", KEYS[1])
else
    return 0
end
```

The "lock validity time" is the time we use as the key's time to live. It is both the auto release time, and the time the client has in order to perform the operation required before another client may be able to acquire the lock again, without technically violating the mutual exclusion guarantee, which is only limited to a given window of time from the moment the lock is acquired.

### The Redlock Algorithm

In the distributed version of the algorithm we assume we have N Redis masters. Those nodes are totally independent, so we don't use replication or any other implicit coordination system. We already described how to acquire and release the lock safely in a single instance.

In order to acquire the lock, the client performs the following operations:

1. It gets the current time in milliseconds
2. It tries to acquire the lock in all the N instances sequentially, using the same key name and random value in all the instances. During step 2, when setting the lock in each instance, the client uses a timeout which is small compared to the total lock auto-release time in order to acquire it. For example if the auto-release is 10 seconds, the timeout could be in the ~ 5-50 milliseconds range. This prevents the client from remaining blocked for a long time trying to talk with a Redis node which is down: if an instance is not available, we should try to talk with the next instance ASAP. 
3. The client computes how much time elapsed in order to acquire the lock, by subtracting from the current time the timestamp obtained in step 1. If and only if the client was able to acquire the lock in the majority of the instances (at least 3), and the total time elapsed to acquire the lock is less than lock validity time, the lock is considered to be acquired.
4. If the lock was acquired, its validity time is considered to be initial validity time minus the time elapsed.
5. If the client failed to acquire the lock for some reason (either it was not able to lock N/2 + 1 instances or the validity time is negative), it will try to unlock all the instances (even the instances it belived it was not able to lock).

## What is Distributed Locking?

Distributed Locking is a technique used in distributed systems to prevent different processes from accessing or changing shared data simultaneously. This is crucial for maintaining system consistency, order and preventing conflicts in a distributed system environments.

## Functionality and Features

Distributed Locking provides a set of failures that manage concurrent data access across different nodes. The key features include:

- Exclusive Access: Only one process can hold a lock at any given time.
- Fail-safe: The system automatically releases locks in case of node failures.
- Deadlock Prevention: The system provides mechanisms to avoid or resolve deadlock situations.

## Architecture

A typical distributed locking system consists of multiple nodes and a lock manager. The lock manager maintains the status of locks and handles lock and release requests from nodes. Nodes communicate with the lock manager to gain access to shared resources.

## Distributed Locking In Go

Distributed locking is a crucial concept in distributed systems. It involves the blocking of processes or threads attempting to access a shared resource that is currently in use by another thread or process.

In certain sencarios, it becomes crucial to safeguard shared resources from multiple processes trying to access them simultaneously. This requirement arises when tasks within a system demand a synchronized processing approach. Without such synchronization conflicts and data corruption may occur.  

To address this challenge, synchronization mechanisms are employed. One common method is using locks or mutexes, which act as way to guard shared resources.

Synchronous processing of critical tasks is particularly crucial in areas like concurrent programming, real-time systems, and database management, where data integrity and responsiveness are critical. 

## Distributed Locking in Halodoc

```go
func PayOrders(request views.PayOrdersRequest, lockService lock.BaseLockService) error {
    locker, err := lockService.ObtainLock(request.MerchantLocationId, lock.LockParams{
        TimeToLive: 20,
        RetryInterval: 2,
    })
    if err != nil {
        return err
    }

    _, err := o.DistributedSettlementPayment(requestr)
    lockService.Unlock(locker)

    return nil
}
```

## Summary

This blog post discusses the concept of distributed locking in Go, which ensures that only one process or thread can access a shared resource at a time in distributed systems. The post explains the need for distributed locking and provides an example of implementing it using the sync.Mutex type in Go. However, it highlights that sync.Mutex is not suitable for distributed systems with multiple instances running concurrently and introduces the use of external storage systems like Redis for distributed locking. The post mentions the redislock library as an example and demonstrates its usage. It further explains how the concept of distributed locking is implemented in the common library at Halodoc, providing an interface for distributed locking and an implementation using the redislock library. The post emphasized the benefits of abstracting the internal implementation from the client, allowing for easy swapping of different distributed locking libraries without affecting the client code.

