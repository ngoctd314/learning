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


9 8 7 6 5 4 3 2 1

0 1 2 3 4 5 6 7 8

8 7 4 2 2 1

1 3 5
