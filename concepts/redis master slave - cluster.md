# Redis master-slave and Redis clusters and the difference between them?

## Redis master slave

In this mode, there are a set of nodes that are masters and a set of nodes that are slaves or replicas. They are generally in pairs and slaves replicate the masters.

The client and read and write to masters and only read from slaves. The other masters are in sync with this one so that in case of failure of the main master the other ones can take its place and start serving. This is generally not used as it is a very costly setup.

Instead when people use are 1 master and multiple slaves and once the master goes down the slave is promoted as master and the master will be changed to a slave once it comes back. 

### How does failover work?

There can be two scenarios either a master going down or a shard going down.

**Master Failure**

In this scenarios, there can be two cases. If you have hot standby master you can promote it to master and the replicas of the older master will be converted to master but now it is a hot standby.

Another case is your replicas will be promoted to master and once the old node is fine it will take up the role of the replicas.

**Slave Failover**

This case is very simple. There is no failover here. There will be just slave reboot and partial or full sync with the master.

## Redis cluter

Clustering is different that master-slave in a very basic sense that not all the data of the cluster resides on one node. There are multiple master nodes where data will reside and each master has a different group of data which is called a shard. There are 16384 shards divided among the total number of masters. Each master has replicas and when the master goes down the replicas takes over as master.

Cluster is very efficient on a high scale especially when the writes are your problem.

### How failover works in a cluster

In the case of cluster whenever a node fails it again can be of two types master or replica. In the case of master failure the replica will take up the role of master and the cluster will keep working.

In case of replica failure, there is no impact the slave will reboot and them does partial or full sync and then again start working as a replica.

## When to use master-slave or clustering?

Whenever there is amount of data that cannot fit your ram you can shard it and save it in the cluster. So when the amount of data is more than ram of single-node use cluster. Otherwise, you can use master-slave. Also, note that sharding can also help to scale your writes on Redis.
