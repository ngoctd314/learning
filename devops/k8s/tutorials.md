## What is K8s

- Open source container orchestration tool
- Developed by Google
- Helps you manage containerized applications in different deployment environments

## What features do orchestration tools offer?

- High availability or no downtime
- Scalability or high performance
- Disaster recovery back up and restore

## Architecture

Basic architecture contains a master node and some workers node.

The kubelet is the primary "node agent" that runs on each node. It can register the node with the apiserver using one of: the hostname, a flag to override the hostname; or specific logic for a cloud provider.

**Api Server**

The Kubernetes API server validates and configures data for the api objects which include pods, services, replicationcontrolelrs, and other. The API server services REST operations and provides the frontend to the cluster's shared state through which all other components interact.

**Controller manager**

The k8s controller manager is a daemon that embeds the core control loops shipped with k8s. In applications of robotics and automation, a control loop is a non-terminating loop that regulates the state of the system.

**Scheduler**

In k8s, scheduling refers to making sure that Pods are matched to Nodes so that Kubelet can run them.

A scheduler watches for newly created Pods that have no Node assigned. For every Pod that scheduler discovers, the scheduler becoms responsible for finding the best Node for that Pod to run on.

### Pods

Pods are the smallest deployable units of computing that you can create and manage in K8S.
