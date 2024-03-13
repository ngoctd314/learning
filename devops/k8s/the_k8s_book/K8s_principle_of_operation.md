# Kubernetes principles of operation

## K8s from 40K feet

At the highest level, K8s is two things:

- A cluster for running applications
- An orchestrator of cloud-native microservices apps

## K8s as a cluster

K8s is like any other cluster - a bunch of machines to hose applications. We call these machines "nodes" and they can by physical servers, virtual machines, cloud instances, Raspberry Pis, and more.

A k8s cluster is made of a control plane and nodes. The control plane exposes the API, has a scheduler for assigning work, and records the state of the cluster and apps in persistent store. Nodes are where user applications run.

It can be useful to think of the control plane as the brains of the cluster, and the nodes as the muscle. In the analogy, the control plane is the brains because it implements the clever features such as scheduling, auto-scaling, and zero-downtime rolling updates. The nodes are the mu


## K8s as an orchestrator

Orchestrator is just a fancy word for a system that takes care of deploying and managing applications.

In the sports world we call this coaching. In the application world we call it orchestration. Kubernetes orchestrates cloud-native microservices applications.

## How it works

To make this happen, you start out with an app, package it up and give it to the cluster (Kubernetes). The cluster is made up of one or more masters and a bunch of nodes.

## K8s as an orchestrator

Orchestrator is just a fancy word for a system that takes care of deploying and managing applications.

## How it work

You start out with an app 

## Control plane and worker nodes

As previously mentioned, a Kubernetes cluster is made of control plane nodes and worker nodes. These are Linux hosts that can be VM, bare metal servers in your datacenter, or instances in a private or public cloud.

**The control plane**

A K8s control plane node is a server running collection of system services that make up the control plane of the cluster. Sometimes we call them Masters, Heads or Head nodes.

The simplest setups run a single control plane node. However, this is only suitable for labs and test environments. For production environments, multiple control plane nodes configured for high availability (HA) is vital. Generally speaking, 3 or 5 is recommended for HA.

Let's take a look at the different services making up the control plane.

**The API server**

The API server is the Grand Central of K8s. All communication, between all components, must go through the API server. We'll get into the detail later, but it's important to understand that internal system components, as well as external user componennts, all communicate via the API server - all roads lead to the API server.

It exposes a RESTful API that you POST YAML configuration files to over HTTPS. These YAML files which sometimes call manifests, describe the desired state of an application.

All requests to the API server are subject to authentication and authorization checks. Once these are done, the config in the YAML file is validated, persisted to the cluster store, and work is scheduled to the cluster.

**The cluster store**

The cluster store is the only stateful part of the control plane and persistently stores the entire configuration and state of the cluster. As such, it's a vital component of every K8s cluster - no cluster store, no cluster.

The cluster store is currently based on etcd, a popular distributed database. As it's the single source of truth for a cluster, you should run between 3-5 etcd replicas for HA, and you should provide adequate ways to recover when things go wrong.

On the topic of availability, etcd prefer consistency over availability. This means it doesn't tolerate split-brains and will halt updates to the cluster in order to maintain consistency. However, if this happens, use applications should continue to work, you just won't be able to update the cluster config.

**The controller manager and controllers**

Some of the controllers include the Deployment controller, the StatefulSet controller, and the ReplicaSet controller.

**The Scheduler**

At a high level, the scheduler watches the API server for new work tasks and assigns them to appropriate healthy worker nodes. Behind the scenes, it implements complex logic that filters out nodes incapable of running tasks.

If the scheduler doesn't find a suitable node, the task isn't scheduled and gets marked as pending.

The scheduler isn't reponsible for running tasks, just picking the nodes to run them. A task is normally a Pod/container. You'll learn about Pods and containers in later.

**The cloud controller manager**


